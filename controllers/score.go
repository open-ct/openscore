package controllers

import (
	"encoding/json"
	"log"
	"math"
	"openscore/models"
	"strconv"
	"strings"
	"time"
)

func (c *TestPaperApiController) Display() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	log.Println(requestBody["testId"])
	testIdstr := requestBody["testId"].(string)

	testId, err := strconv.ParseInt(testIdstr, 10, 64)
	if err != nil {
		log.Println("parse questionId fail")
	}

	var testPaper models.TestPaper
	var topic models.Topic
	var subTopic []models.SubTopic
	testPaper.GetTestPaper(testId)
	topic.GetTopic(testPaper.Question_id)
	models.GetSubTopicsByTestId(testPaper.Question_id, &subTopic)
	type subTopicRes struct {
		Question_detail_id    int64
		Question_detail_name  string
		Question_id           int64
		Question_detail_score int64
		Test_detail_id        int64
	}
	var testInfoList []models.TestPaperInfo
	var subTopics []subTopicRes
	for i := 0; i < len(subTopic); i++ {
		var testPaperInfo models.TestPaperInfo
		testPaperInfo.GetTestPaperInfoByTestIdAndQuestionDetailId(testId, subTopic[i].Question_detail_id)
		tempTopic := subTopicRes{subTopic[i].Question_detail_id, subTopic[i].Question_detail_name, subTopic[i].Question_id, subTopic[i].Question_detail_score, (testPaperInfo.Test_detail_id)}
		subTopics = append(subTopics, tempTopic)
		log.Println(subTopics)
		testInfoList = append(testInfoList, testPaperInfo)
	}
	data := make(map[string]interface{})
	data["questionId"] = testPaper.Question_id

	data["questionName"] = topic.Question_name
	data["subTopic"] = subTopics
	data["picSrcs"] = testInfoList
	resp := Response{"10000", "OK", data}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) List() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)

	userIdstr := requestBody["userId"].(string)

	userId, err := strconv.ParseInt(userIdstr, 10, 64)
	if err != nil {
		log.Println("parse userId fail")
	}
	var papers []models.UnderCorrectedPaper
	models.GetDistributedPaperByUserId(userId, &papers)
	data := make(map[string]interface{})
	data["papers"] = papers
	resp := Response{"10000", "OK", data}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) Point() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	userIdstr := requestBody["userId"].(string)
	scoresstr := requestBody["scores"].(string)
	testIdstr := requestBody["testId"].(string)
	testDetailIdstr := requestBody["testDetailId"].(string)
	// userId, _ := strconv.ParseInt(userIdstr, 10, 64)
	userId := userIdstr
	scores := strings.Split(scoresstr, "-")
	testDetailIds := strings.Split(testDetailIdstr, "-")
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	var scoreArr []int64
	var sum int64 = 0
	for _, i := range scores {
		j, err := strconv.ParseInt(i, 10, 64)
		sum += j
		if err != nil {
			panic(err)
		}
		scoreArr = append(scoreArr, j)
	}

	var test models.TestPaper
	var topic models.Topic
	test.GetTestPaper(testId)
	topic.GetTopic(test.Question_id)
	// var testInfos []models.TestPaperInfo
	// models.GetTestInfoListByTestId(testId, &testInfos)

	var underTest models.UnderCorrectedPaper
	underTest.GetUnderCorrectedPaper(userId, testId)
	// underTest.Delete()

	final := false

	if topic.Score_type == 1 {
		test.Examiner_first_id = userId
		test.Examiner_first_score = sum
		final = true
	} else if topic.Score_type == 2 && test.Examiner_first_id == "-1" {
		test.Examiner_first_id = userId
		test.Examiner_first_score = sum
	} else if topic.Score_type == 2 && test.Examiner_second_id == "-1" {
		test.Examiner_second_id = userId
		test.Examiner_second_score = sum
		if math.Abs(float64(test.Examiner_second_score)-float64(test.Examiner_first_score)) <= float64(topic.Standard_error) {
			log.Println(math.Abs(float64(test.Examiner_second_score) - float64(test.Examiner_first_score)))
			sum = int64(math.Abs(float64(test.Examiner_second_score+test.Examiner_first_score)) / 2)
			log.Println("hello world")
			final = true
		} else {
			newUnderTest := models.UnderCorrectedPaper{}
			newUnderTest.User_id = "10000"
			newUnderTest.Test_question_type = 3
			newUnderTest.Test_id = underTest.Test_id
			newUnderTest.Question_id = underTest.Question_id
			newUnderTest.Save()
		}
	}
	if underTest.Test_question_type == 4 || underTest.Test_question_type == 5 {
		test.Leader_id = userId
		test.Leader_score = sum
		final = true
	} else if underTest.Test_question_type == 3 {
		test.Examiner_third_id = userId
		test.Examiner_third_score = sum
		first := math.Abs(float64(test.Examiner_third_score - test.Examiner_first_score))
		second := math.Abs(float64(test.Examiner_third_score - test.Examiner_second_score))
		var small float64
		if first <= second {
			small = first
			sum = (test.Examiner_third_score + test.Examiner_first_score) / 2
		} else {
			small = second
			sum = (test.Examiner_third_score + test.Examiner_second_score) / 2
		}
		if small <= float64(topic.Standard_error) {
			// test.Final_score = sum
			final = true
		} else {
			test.Question_status = 2

			newUnderTest := models.UnderCorrectedPaper{}
			newUnderTest.User_id = "10000"
			newUnderTest.Test_question_type = 4
			newUnderTest.Test_id = underTest.Test_id
			newUnderTest.Question_id = underTest.Question_id
			newUnderTest.Save()

		}
		//??
	}
	if final {
		//???
		test.Final_score = sum
	}
	//  else {
	// 	newUnderTest := underTest
	// 	newUnderTest.User_id = 10000
	// 	// newUnderTest.Test_question_type += 1
	// 	newUnderTest.Save()
	// }
	underTest.Delete()
	test.Update()
	for i := 0; i < len(scores); i++ {
		score := scoreArr[i]
		var tempTest models.TestPaperInfo
		id, _ := strconv.ParseInt(testDetailIds[i], 10, 64)
		log.Println(id)
		tempTest.GetTestPaperInfo(id)
		if topic.Score_type == 1 {
			tempTest.Examiner_first_id = userId
			tempTest.Examiner_first_score = score
		} else if topic.Score_type == 2 && tempTest.Examiner_first_id == "-1" {
			tempTest.Examiner_first_id = userId
			tempTest.Examiner_first_score = score
		} else if topic.Score_type == 2 && tempTest.Examiner_second_id == "-1" {
			tempTest.Examiner_second_id = userId
			tempTest.Examiner_second_score = score
			// if final{
			// 	score =  int64(math.Abs(float64(tempTest.Examiner_second_score+tempTest.Examiner_first_score)) / 2)
			// }
		}
		if underTest.Test_question_type == 4 || underTest.Test_question_type == 5 {
			tempTest.Leader_id = userId
			tempTest.Leader_score = score
		} else if underTest.Test_question_type == 3 {
			tempTest.Examiner_third_id = userId
			tempTest.Examiner_third_score = score
		}
		if final {
			tempTest.Final_score = score
		}
		tempTest.Update()
	}

	var record models.ScoreRecord
	record.Score = sum
	record.Question_id = topic.Question_id
	record.Test_id = testId
	record.Test_record_type = underTest.Test_question_type
	record.User_id = userId
	record.Score_time = time.Now()
	record.Save()
}

func (c *TestPaperApiController) Problem() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	userIdstr := requestBody["userId"].(string)
	problemTypestr := requestBody["problemType"].(string)
	testIdstr := requestBody["testId"].(string)
	// userId, _ := strconv.ParseInt(userIdstr, 10, 64)
	userId := userIdstr
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	problemType, _ := strconv.ParseInt(problemTypestr, 10, 64)
	var underTest models.UnderCorrectedPaper
	var record models.ScoreRecord
	var test models.TestPaper

	underTest.GetUnderCorrectedPaper(userId, testId)
	var newUnderTest = underTest
	underTest.Delete()
	newUnderTest.User_id = "10000"
	newUnderTest.Test_question_type = 6
	newUnderTest.Problem_type = problemType
	has, _ := newUnderTest.IsDuplicate()
	if !has {
		log.Println("dup")
		newUnderTest.Save()
		test.GetTestPaper(testId)
		test.Question_status = 3
		test.Update()
	}

	record.Test_record_type = 5
	record.Test_id = testId
	record.User_id = userId
	record.Question_id = test.Question_id
	record.Test_record_type = 5
	record.Save()
}

func (c *TestPaperApiController) Answer() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	// userIdstr := requestBody["userId"].(string)
	testIdstr := requestBody["testId"].(string)
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	var test models.TestPaper
	test.GetTestPaper(testId)
	var answerTest models.TestPaper
	answerTest.GetTestPaperByQuestionIdAndQuestionStatus(test.Question_id, 5)

	var as []models.TestPaperInfo
	models.GetTestInfoListByTestId(answerTest.Test_id, &as)
	data := make(map[string]interface{})
	data["keyTest"] = as
	resp := Response{"10000", "ok", data}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) ExampleDeatil() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	// userIdstr := requestBody["userId"].(string)
	testIdstr := requestBody["exampleTestId"].(string)
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	var test models.TestPaper
	test.GetTestPaper(testId)
	var exampleTest []models.TestPaper
	//??
	models.GetTestPaperListByQuestionIdAndQuestionStatus(test.Question_id, 6, &exampleTest)

	var topic models.Topic
	topic.GetTopic(exampleTest[0].Question_id)
	var tests [][]models.TestPaperInfo
	for i := 0; i < len(exampleTest); i++ {
		var temp []models.TestPaperInfo
		models.GetTestInfoListByTestId(exampleTest[i].Test_id, &temp)
		tests = append(tests, temp)
	}
	data := make(map[string]interface{})
	data["questionName"] = topic.Question_name
	data["test"] = tests
	resp := Response{"10000", "ok", data}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) ExampleList() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	// userIdstr := requestBody["userId"].(string)
	testIdstr := requestBody["testId"].(string)
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	var testPaper models.TestPaper
	testPaper.GetTestPaper(testId)
	var exampleTest []models.TestPaper
	//??
	models.GetTestPaperListByQuestionIdAndQuestionStatus(testPaper.Question_id, 6, &exampleTest)
	data := make(map[string]interface{})
	data["exampleTestId"] = exampleTest
	resp := Response{"10000", "ok", data}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) Review() {
	defer c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	userIdstr := requestBody["userId"].(string)
	// userId, _ := strconv.ParseInt(userIdstr, 10, 64)
	userId := userIdstr
	var records []models.ScoreRecord
	models.GetLatestRecores(userId, &records)
	data := make(map[string]interface{})
	data["records"] = records
	resp := Response{"10000", "ok", data}
	c.Data["json"] = resp
}

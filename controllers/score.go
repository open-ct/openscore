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
	// topic.GetTopic(testPaper.Question_id)
	models.GetSubTopicsByTestId(testPaper.Test_id, &subTopic)
	var picSrcs []string
	for i := 0; i < len(subTopic); i++ {
		var testPaperInfo models.TestPaperInfo
		testPaperInfo.GetTestPaperInfoByTestIdAndQuestionDetailId(subTopic[i].Question_id, subTopic[i].Question_detail_id)
		picSrcs = append(picSrcs, testPaperInfo.Pic_src)
	}
	data := make(map[string]interface{})
	data["questionId"] = testPaper.Question_id

	data["questionName"] = topic.Question_name
	data["subTopic"] = subTopic
	data["picSrcs"] = picSrcs
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
	userId, _ := strconv.ParseInt(userIdstr, 10, 64)
	scores := strings.Split(scoresstr, "-")
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
	var testInfos []models.TestPaperInfo
	models.GetTestInfoListByTestId(testId, &testInfos)

	var underTest models.UnderCorrectedPaper
	underTest.GetUnderCorrectedPaper(testId)
	underTest.Delete()

	final := false

	if underTest.Test_question_type == 1 {
		test.Examiner_first_id = userId
		test.Examiner_first_score = sum
		final = true
	} else if underTest.Test_question_type == 2 && test.Examiner_first_id == -1 {
		test.Examiner_first_id = userId
		test.Examiner_first_score = sum
	} else if underTest.Test_question_type == 2 && test.Examiner_second_id == -1 {
		test.Examiner_second_id = userId
		test.Examiner_second_score = sum
		if math.Abs(float64(test.Examiner_second_score)-float64(test.Examiner_first_score)) <= float64(topic.Standard_error) {
			final = true
		}
	} else if underTest.Test_question_type == 4 || underTest.Test_question_type == 5 {
		test.Leader_id = userId
		test.Leader_score = sum
		final = true
	} else {
		test.Examiner_third_id = userId
		test.Examiner_third_score = sum
	}
	if final {
		test.Final_score = sum
		underTest.Delete()
	} else {
		newUnderTest := underTest
		underTest.Delete()
		newUnderTest.Test_question_type += 1
		newUnderTest.Save()
	}
	test.Update()
	for i := 0; i < len(scores); i++ {
		score, _ := strconv.ParseInt(scores[i], 10, 64)
		sum += score
		if underTest.Test_question_type == 1 {
			testInfos[i].Examiner_first_id = userId
			testInfos[i].Examiner_first_score = score
		} else if underTest.Test_question_type == 2 && testInfos[i].Examiner_first_id == -1 {
			testInfos[i].Examiner_first_id = userId
			testInfos[i].Examiner_first_score = score
		} else if underTest.Test_question_type == 2 && testInfos[i].Examiner_second_id == -1 {
			testInfos[i].Examiner_second_id = userId
			testInfos[i].Examiner_second_score = score
		} else if underTest.Test_question_type == 4 || underTest.Test_question_type == 5 {
			testInfos[i].Leader_id = userId
			testInfos[i].Leader_score = score
		} else {
			testInfos[i].Examiner_second_id = userId
			testInfos[i].Examiner_second_score = score
		}
		if final {
			testInfos[i].Final_score = score
		}
	}

	var record models.ScoreRecord
	record.Score = sum
	record.Test_id = testId
	record.Test_record_type = 1
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
	userId, _ := strconv.ParseInt(userIdstr, 10, 64)
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	problemType, _ := strconv.ParseInt(problemTypestr, 10, 64)
	var underTest models.UnderCorrectedPaper
	var record models.ScoreRecord
	var test models.TestPaper
	test.GetTestPaper(testId)
	test.Problem_type = problemType
	test.Update()
	record.Test_record_type = 5
	record.Test_id = testId
	record.User_id = userId
	record.Question_id = test.Question_id
	record.Test_record_type = 5
	record.Save()
	underTest.GetUnderCorrectedPaper(testId)
	var newUnderTest = underTest
	underTest.Delete()
	newUnderTest.Test_question_type = 5
	newUnderTest.Save()
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
	answerTest.GetTestPaperByQuestionIdAndQuestionStatus(test.Question_id, 3)

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
	var topic models.Topic
	topic.GetTopic(test.Question_id)
	var tests []models.TestPaperInfo
	models.GetTestInfoListByTestId(testId, &tests)
	data := make(map[string]interface{})
	data["questionName"] = topic.Question_name
	data["test"] = tests
	resp := Response{"10000", "ok", data}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) ExampleList() {
	c.ServeJSON()
	var requestBody map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	// userIdstr := requestBody["userId"].(string)
	testIdstr := requestBody["testId"].(string)
	testId, _ := strconv.ParseInt(testIdstr, 10, 64)
	var testPaper models.TestPaper
	testPaper.GetTestPaper(testId)
	data := make(map[string]interface{})
	data["exampleTestId"] = testPaper.Example_test_id
	resp := Response{"10000", "ok", data}
	c.Data["json"] = resp

}

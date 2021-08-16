package controllers

import (
	"encoding/json"
	"log"
	"math"
	"openscore/models"
	"openscore/requests"
	"openscore/responses"
	"strconv"
	"strings"
	"time"
)

func (c *TestPaperApiController) Display() {
	defer c.ServeJSON()
	var requestBody requests.TestDisplay
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.TestId

	var testPaper models.TestPaper
	var topic models.Topic
	var subTopic []models.SubTopic
	var response responses.TestDisplay

	err = testPaper.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	err = topic.GetTopic(testPaper.Question_id)
	if err != nil {
		resp := Response{"10003", "get topic fail", err}
		c.Data["json"] = resp
		return
	}
	err = models.GetSubTopicsByTestId(testPaper.Question_id, &subTopic)
	if err != nil {
		resp := Response{"10004", "get subtopic fail", err}
		c.Data["json"] = resp
		return
	}

	for i := 0; i < len(subTopic); i++ {
		var testPaperInfo models.TestPaperInfo
		err = testPaperInfo.GetTestPaperInfoByTestIdAndQuestionDetailId(testId, subTopic[i].Question_detail_id)
		if err != nil {
			resp := Response{"10005", "get testPaperInfo fail", err}
			c.Data["json"] = resp
			return
		}
		tempSubTopic := responses.SubTopicPlus{SubTopic: subTopic[i], Test_detail_id: testPaperInfo.Test_detail_id}

		response.SubTopics = append(response.SubTopics, tempSubTopic)
		response.TestInfos = append(response.TestInfos, testPaperInfo)
	}
	response.QuestionId = topic.Question_id
	response.QuestionName = topic.Question_name
	response.TestId = testId
	resp := Response{"10000", "OK", response}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) List() {
	defer c.ServeJSON()
	var requestBody requests.TetsList
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	log.Println(requestBody)

	userId := requestBody.UserId

	var response responses.TestList
	err = models.GetDistributedTestIdPaperByUserId(userId, &response.TestId)
	if err != nil {
		resp := Response{"10002", "get distribution fail", err}
		c.Data["json"] = resp
		return
	}
	if len(response.TestId) == 0 {
		resp := Response{"10003", "there is no paper to correct", err}
		c.Data["json"] = resp
		return

	}
	log.Println(response)
	resp := Response{"10000", "OK", response}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) Point() {
	defer c.ServeJSON()
	var requestBody requests.TestPoint
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	log.Println(requestBody)
	userId := requestBody.UserId
	scoresstr := requestBody.Scores
	testId := requestBody.TestId
	testDetailIdstr := requestBody.TestDetailId
	scores := strings.Split(scoresstr, "-")
	testDetailIds := strings.Split(testDetailIdstr, "-")
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
	err = test.GetTestPaper(testId)
	if err != nil || test.Test_id == 0 {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	err = topic.GetTopic(test.Question_id)
	if err != nil || topic.Question_id == 0 {
		resp := Response{"10003", "get topic fail", err}
		c.Data["json"] = resp
		return
	}

	var underTest models.UnderCorrectedPaper
	err = underTest.GetUnderCorrectedPaper(userId, testId)
	if err != nil || underTest.Question_id == 0 {
		resp := Response{"10004", "get underCorrected fail", err}
		c.Data["json"] = resp
		return
	}

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
			final = true
		} else {
			newUnderTest := models.UnderCorrectedPaper{}
			newUnderTest.User_id = "10000"
			newUnderTest.Test_question_type = 3
			newUnderTest.Test_id = underTest.Test_id
			newUnderTest.Question_id = underTest.Question_id
			err = newUnderTest.Save()
			if err != nil {
				resp := Response{"10005", "insert undertest fail", err}
				c.Data["json"] = resp
				return
			}
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
			final = true
		} else {
			test.Question_status = 2

			newUnderTest := models.UnderCorrectedPaper{}
			newUnderTest.User_id = "10000"
			newUnderTest.Test_question_type = 4
			newUnderTest.Test_id = underTest.Test_id
			newUnderTest.Question_id = underTest.Question_id
			err = newUnderTest.Save()
			if err != nil {
				resp := Response{"10006", "insert undertest fail", err}
				c.Data["json"] = resp
				return
			}
		}
	}
	if final {
		test.Final_score = sum
	}
	err = underTest.Delete()
	if err != nil {
		resp := Response{"10006", "delete undertest fail", err}
		c.Data["json"] = resp
		return
	}
	err = test.Update()
	if err != nil {
		resp := Response{"10007", "update test fail", err}
		c.Data["json"] = resp
		return
	}
	for i := 0; i < len(scores); i++ {
		score := scoreArr[i]
		var tempTest models.TestPaperInfo
		id, _ := strconv.ParseInt(testDetailIds[i], 10, 64)
		log.Println(id)
		err = tempTest.GetTestPaperInfo(id)
		if err != nil {
			resp := Response{"10008", "get testPaper fail", err}
			c.Data["json"] = resp
			return
		}
		if topic.Score_type == 1 {
			tempTest.Examiner_first_id = userId
			tempTest.Examiner_first_score = score
		} else if topic.Score_type == 2 && tempTest.Examiner_first_id == "-1" {
			tempTest.Examiner_first_id = userId
			tempTest.Examiner_first_score = score
		} else if topic.Score_type == 2 && tempTest.Examiner_second_id == "-1" {
			tempTest.Examiner_second_id = userId
			tempTest.Examiner_second_score = score
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
		err = tempTest.Update()
		if err != nil {
			resp := Response{"10009", "update testPaper fail", err}
			c.Data["json"] = resp
			return
		}
	}

	var record models.ScoreRecord
	record.Score = sum
	record.Question_id = topic.Question_id
	record.Test_id = testId
	record.Test_record_type = underTest.Test_question_type
	record.User_id = userId
	record.Score_time = time.Now()
	err = record.Save()
	if err != nil {
		resp := Response{"10010", "insert record fail", err}
		c.Data["json"] = resp
		return
	}
	resp := Response{"10000", "ok", err}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) Problem() {
	defer c.ServeJSON()
	// var requestBody map[string]interface{}
	var requestBody requests.TestProblem
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}

	userId := requestBody.UserId
	problemType := requestBody.ProblemType
	testId := requestBody.TestId

	var underTest models.UnderCorrectedPaper
	var record models.ScoreRecord
	var test models.TestPaper

	err = underTest.GetUnderCorrectedPaper(userId, testId)
	if err != nil {
		resp := Response{"10002", "get underCorrected fail", err}
		c.Data["json"] = resp
		return
	}
	var newUnderTest = underTest
	err = underTest.Delete()
	if err != nil {
		resp := Response{"10002", "delete underTest fail", err}
		c.Data["json"] = resp
		return
	}
	newUnderTest.User_id = "10000"
	newUnderTest.Test_question_type = 6
	newUnderTest.Problem_type = problemType
	has, _ := newUnderTest.IsDuplicate()
	if !has {
		err = newUnderTest.Save()
		if err != nil {
			resp := Response{"10003", "update underTest fail", err}
			c.Data["json"] = resp
			return
		}
		err = test.GetTestPaper(testId)
		if err != nil {
			resp := Response{"10004", "get testPaper fail", err}
			c.Data["json"] = resp
			return
		}
		test.Question_status = 3
		err = test.Update()
		if err != nil {
			resp := Response{"10005", "update testPaper fail", err}
			c.Data["json"] = resp
			return
		}
	}

	record.Test_record_type = 5
	record.Test_id = testId
	record.User_id = userId
	record.Question_id = test.Question_id
	record.Test_record_type = 5
	err = record.Save()
	if err != nil {
		resp := Response{"10006", "insert record fail", err}
		c.Data["json"] = resp
		return
	}
	resp := Response{"10000", "ok", err}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) Answer() {
	defer c.ServeJSON()
	var requestBody requests.TestAnswer
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.TestId
	var test models.TestPaper
	err = test.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	var answerTest models.TestPaper
	err = answerTest.GetTestPaperByQuestionIdAndQuestionStatus(test.Question_id, 5)
	if err != nil {
		resp := Response{"10003", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}

	var as responses.TestAnswer
	err = models.GetTestInfoPicListByTestId(answerTest.Test_id, &as.Pic_src)
	if err != nil {
		resp := Response{"10004", "get testPaperInfo fail", err}
		c.Data["json"] = resp
		return
	}
	resp := Response{"10000", "ok", as}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) ExampleDeatil() {
	defer c.ServeJSON()
	var requestBody requests.ExampleDetail
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.ExampleTestId
	log.Println(testId)
	var test models.TestPaper
	err = test.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	var exampleTest []models.TestPaper
	//??
	err = models.GetTestPaperListByQuestionIdAndQuestionStatus(test.Question_id, 6, &exampleTest)
	if err != nil {
		resp := Response{"10003", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	if len(exampleTest) == 0 {
		resp := Response{"10004", "there is no exampleTest", err}
		c.Data["json"] = resp
		return

	}

	var topic models.Topic
	err = topic.GetTopic(exampleTest[0].Question_id)
	if err != nil {
		resp := Response{"10005", "get topic fail", err}
		c.Data["json"] = resp
		return
	}
	var response responses.ExampleDeatil
	response.QuestionName = topic.Question_name
	for i := 0; i < len(exampleTest); i++ {
		var temp []models.TestPaperInfo
		err = models.GetTestInfoListByTestId(exampleTest[i].Test_id, &temp)
		if err != nil {
			resp := Response{"10006", "get testPaperInfo fail", err}
			c.Data["json"] = resp
			return
		}
		response.Test = append(response.Test, temp)
	}
	resp := Response{"10000", "ok", response}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) ExampleList() {
	defer c.ServeJSON()
	var requestBody requests.ExampleList
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.TestId
	var testPaper models.TestPaper
	err = testPaper.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	var response responses.ExampleList
	err = models.GetTestPaperListByQuestionIdAndQuestionStatus(testPaper.Question_id, 6, &response.TestPapers)
	if err != nil {
		resp := Response{"10003", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	resp := Response{"10000", "ok", response}
	c.Data["json"] = resp

}

func (c *TestPaperApiController) Review() {
	defer c.ServeJSON()
	var requestBody requests.TestReview
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	userId := requestBody.UserId
	var records []models.ScoreRecord
	var response responses.TestReview
	err = models.GetLatestRecords(userId, &records)
	if err != nil {
		resp := Response{"10002", "get record fail", err}
		c.Data["json"] = resp
		return
	}
	for i := 0; i < len(records); i++ {
		response.TestId = append(response.TestId, records[i].Test_id)
		response.Score = append(response.Score, records[i].Score)
		response.ScoreTime = append(response.ScoreTime, records[i].Score_time)
	}
	resp := Response{"10000", "ok", response}
	c.Data["json"] = resp
}

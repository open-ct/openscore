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
	testIds := strings.Split(testIdstr, "-")
	var testInfo models.TestPaperInfo
	var test models.TestPaper
	var sum int64
	for i := 0; i < len(testIds); i++ {
		var underTest models.UnderCorrectedPaper
		id, _ := strconv.ParseInt(testIds[i], 10, 64)
		testInfo.GetTestPaperInfo(id)
		underTest.GetUnderCorrectedPaper(id)
		underTest.Delete()
		testInfo.Examiner_first_id = userId
		score, _ := strconv.ParseInt(scores[i], 10, 64)
		sum += score
		testInfo.Examiner_first_score = score
		testInfo.Final_score = score
	}
	test.GetTestPaper(testInfo.Test_id)
	test.Examiner_first_id = userId
	test.Examiner_first_score = sum
	test.Final_score = sum
	test.Update()

	var record models.ScoreRecord
	record.Score = sum
	record.Test_id = testInfo.Test_id
	record.Test_record_type = 1
	record.User_id = userId
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
	record.Score_type = 1
	record.Test_id = testId
	record.User_id = userId
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
	keyId := test.Answer_test_id
	var as []models.TestPaperInfo
	models.GetTestInfoListByTestId(keyId, &as)
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

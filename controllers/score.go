package controllers

import (
	"encoding/json"
	"log"
	"openscore/models"
	"strconv"
)

type testStruct struct {
	Name string
}

func (c *TestPaperApiController) Display() {
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
	models.GetSubTopicsByQuestionId(testPaper.Question_id, &subTopic)
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
	c.ServeJSON()
}

package controllers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math"
	"openscore/models"
	"openscore/requests"
	"openscore/responses"
	"os"
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

	_, err = testPaper.GetTestPaper(testId)
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
		picName := testPaperInfo.Pic_src
		//图片地址拼接 ，按服务器
		// src := "C:\\Users\\yang\\Desktop\\阅卷系统\\img\\" + picName
		src := "./img/" + picName
		bytes, err := os.ReadFile(src)
		if err != nil {
			log.Println(err)
			resp := Response{"30020", "get 图片显示异常 ", err}
			c.Data["json"] = resp
			return
		}
		encoding := base64.StdEncoding.EncodeToString(bytes)
		tempTestPaperInfo := responses.TestPaperInfoPlus{TestPaperInfo: testPaperInfo, PicCode: encoding}
		response.TestInfos = append(response.TestInfos, tempTestPaperInfo)
	}
	response.QuestionId = topic.Question_id
	response.QuestionName = topic.Question_name
	response.TestId = testId
	resp := Response{"10000", "OK", response}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) List() {
	defer c.ServeJSON()
	var requestBody requests.TestList
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	log.Println(requestBody)
	userId := requestBody.UserId
	var response responses.TestList
	//----------------------------------------------------

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
	scoresStr := requestBody.Scores
	testId := requestBody.TestId
	testDetailIdStr := requestBody.TestDetailId
	scores := strings.Split(scoresStr, "-")
	testDetailIds := strings.Split(testDetailIdStr, "-")
	//-------------------------------------------------------
	//score数组string转int
	var scoreArr []int64
	var sum int64 = 0
	var record models.ScoreRecord
	for _, i := range scores {
		j, err := strconv.ParseInt(i, 10, 64)
		sum += j
		if err != nil {
			panic(err)
		}
		scoreArr = append(scoreArr, j)
	}
	//获取该试卷大题 和抽象大题信息
	var test models.TestPaper
	var topic models.Topic
	_, err = test.GetTestPaper(testId)
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
	//获取试卷未批改表信息（试卷批改状态类型）
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
	} else if underTest.Test_question_type == 2 && test.Examiner_first_id == "-1" {
		test.Examiner_first_id = userId
		test.Examiner_first_score = sum
	} else if underTest.Test_question_type == 2 && test.Examiner_second_id == "-1" {
		test.Examiner_second_id = userId
		test.Examiner_second_score = sum
		if math.Abs(float64(test.Examiner_second_score)-float64(test.Examiner_first_score)) <= float64(topic.Standard_error) {
			log.Println(math.Abs(float64(test.Examiner_second_score) - float64(test.Examiner_first_score)))
			sum = int64(math.Abs(float64(test.Examiner_second_score+test.Examiner_first_score)) / 2)
			final = true
		} else {
			newUnderTest := models.UnderCorrectedPaper{}
			//随机 抽一个 人

			newUnderTest.User_id = models.FindNewUserId(test.Examiner_first_id, test.Examiner_second_id, test.Question_id)
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
	if underTest.Test_question_type == 0  {

		test.Leader_id = userId
		test.Leader_score = sum
		final = true
	}
	if underTest.Test_question_type == 3 {
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
			//应该可以去掉
			test.Question_status = 2

			newUnderTest := models.UnderCorrectedPaper{}
			//阅卷组长类型默认 id 10000
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
		record.Test_finish=1
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
	problemMessage := requestBody.ProblemMessage

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

	newUnderTest.User_id = userId
	newUnderTest.Test_question_type = 6
	newUnderTest.Problem_type = problemType
	newUnderTest.Problem_message=problemMessage
	has, _ := newUnderTest.IsDuplicate()
	if !has {
		err = newUnderTest.Save()
		if err != nil {
			resp := Response{"10003", "update underTest fail", err}
			c.Data["json"] = resp
			return
		}
		_, err = test.GetTestPaper(testId)
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
	record.Problem_type=problemType
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
	_, err = test.GetTestPaper(testId)
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
	var tempString []string
	err = models.GetTestInfoPicListByTestId(answerTest.Test_id, &tempString)
	if err != nil {
		resp := Response{"10004", "get testPaperInfo fail", err}
		c.Data["json"] = resp
		return
	}
	//改成base64编码
	for i := 0; i < len(tempString); i++ {
		picName := tempString[i]
		//图片地址拼接 ，按服务器
		// src:="C:\\Users\\yang\\Desktop\\阅卷系统\\img\\"+picName
		src := "./img/" + picName
		bytes, err := os.ReadFile(src)
		if err != nil {
			log.Println(err)
			resp := Response{"30020", "get 图片显示异常 ", err}
			c.Data["json"] = resp
			return
		}
		encoding := base64.StdEncoding.EncodeToString(bytes)
		as.Pics = append(as.Pics, encoding)
	}
	resp := Response{"10000", "ok", as}
	c.Data["json"] = resp
}

func (c *TestPaperApiController) ExampleDetail() {
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
	//____________________________________________________________
	var test models.TestPaper
	_, err = test.GetTestPaper(testId)
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
	var response responses.ExampleDetail
	response.QuestionName = topic.Question_name
	for i := 0; i < len(exampleTest); i++ {
		var temp []models.TestPaperInfo
		err = models.GetTestInfoListByTestId(exampleTest[i].Test_id, &temp)
		if err != nil {
			resp := Response{"10006", "get testPaperInfo fail", err}
			c.Data["json"] = resp
			return
		}
		//转64编码
		for j := 0; j < len(temp); j++ {
			picName := temp[j].Pic_src
			// src:="C:\\Users\\yang\\Desktop\\阅卷系统\\img\\"+picName
			src := "./img/" + picName
			bytes, err := os.ReadFile(src)
			if err != nil {
				log.Println(err)
				resp := Response{"30020", "get 图片显示异常 ", err}
				c.Data["json"] = resp
				return
			}
			encoding := base64.StdEncoding.EncodeToString(bytes)
			temp[j].Pic_src = encoding
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
	//----------------------------------------------------------------------
	var testPaper models.TestPaper
	_, err = testPaper.GetTestPaper(testId)
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
func (c *TestPaperApiController) ReviewPoint() {
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
	var record models.ScoreRecord
	for _, i := range scores {
		j, err := strconv.ParseInt(i, 10, 64)
		sum += j
		if err != nil {
			panic(err)
		}
		scoreArr = append(scoreArr, j)
	}

	var test models.TestPaper
	_, err = test.GetTestPaper(testId)
	if err != nil || test.Test_id == 0 {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	//判断是否二次阅卷
	var topic models.Topic
	topic.GetTopic(test.Question_id)
	scoreType := topic.Score_type

	num := 0
	if test.Examiner_first_id == userId {
		num = 0
		test.Examiner_first_score = sum
		if scoreType == 1 {
			test.Final_score = sum
			record.Test_finish=1
		}
		var record models.ScoreRecord
		record.GetRecordByTestId(testId, userId)
		record.Score = sum
		record.Update()

	} else if test.Examiner_second_id == userId {
		num = 1
		test.Examiner_second_score = sum
		var record models.ScoreRecord
		record.GetRecordByTestId(testId, userId)
		record.Score = sum
		record.Update()
	} else {
		num = 2
		test.Examiner_third_score = sum
		var record models.ScoreRecord
		record.GetRecordByTestId(testId, userId)
		record.Score = sum
		record.Update()
	}
	err = test.Update()
	if err != nil || test.Test_id == 0 {
		resp := Response{"10003", "update test paper fail", err}
		c.Data["json"] = resp
		return
	}

	for i := 0; i < len(testDetailIds); i++ {
		var testInfo models.TestPaperInfo
		testInfoId, _ := strconv.ParseInt(testDetailIds[i], 10, 64)
		testInfo.GetTestPaperInfo(testInfoId)
		if num == 0 {
			testInfo.Examiner_first_score = scoreArr[i]
			if scoreType == 1 {
				testInfo.Final_score = scoreArr[i]
			}
		} else if num == 1 {
			testInfo.Examiner_second_score = scoreArr[i]
		} else {
			testInfo.Examiner_third_score = scoreArr[i]
		}
		err = testInfo.Update()
		if err != nil || test.Test_id == 0 {
			resp := Response{"10004", "update testinfo paper fail", err}
			c.Data["json"] = resp
			return
		}
	}
	c.Data["json"] = Response{"10000", "ok", nil}
}
//自评列表 chen
func (c *TestPaperApiController) SelfScoreList() {
	defer c.ServeJSON()
	var requestBody requests.TestList
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	log.Println(requestBody)
	userId := requestBody.UserId
	var response responses.TestList
	//----------------------------------------------------

	err = models.GetUnMarkSelfTestIdPaperByUserId(userId, &response.TestId)
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
/**
20.自评卷打分
*/
func (c *TestPaperApiController) SelfMarkPoint() {
	defer c.ServeJSON()
	var requestBody requests.TestPoint
	var resp Response
	var  err error

	err=json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	userId := requestBody.UserId
	testId := requestBody.TestId
	scoreStr:= requestBody.Scores
	testDetailIdStr:=requestBody.TestDetailId
	testDetailIds := strings.Split(testDetailIdStr, "-")
	scores := strings.Split(scoreStr, "-")

	//---------------------------------------------------------------------------------------


   //查找大题
	var test models.TestPaper
	_,err = test.GetTestPaper(testId)
	if err != nil || test.Test_id == 0 {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	var topic models.Topic
	topic.GetTopic(test.Question_id)
	standardError := topic.Standard_error

	//分三种情况
	if userId == test.Examiner_first_id {
		var sum int64
		//给试卷详情表打分
		for i := 0; i < len(testDetailIds); i++ {
			//取出小题试卷id,和小题分数
			var testInfo models.TestPaperInfo
			testDetailIdString:=testDetailIds[i]
			testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
			scoreString:=scores[i]
			score, _ := strconv.ParseInt(scoreString, 10, 64)
			//------------------------------------------------


			//查试卷小题
			err := testInfo.GetTestPaperInfo(testDetailId)
			if err != nil {
				resp := Response{"10008", "get testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			//修改试卷详情表


			testInfo.Examiner_first_self_score=score

			err = testInfo.Update()
			if err != nil {
				resp := Response{"10009", "update testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			sum += score
		}
		//给试卷表打分

		test.Examiner_first_self_score = sum
		err = test.Update()
		if err != nil {
			resp := Response{"10007", "update test fail", err}
			c.Data["json"] = resp
			return
		}

		//删除试卷待批改表 ，增加试卷记录表
		var record models.ScoreRecord
		var underTest models.UnderCorrectedPaper

		err = models.GetSelfScorePaperByTestQuestionTypeAndTestId(&underTest, testId,userId)
		if err!=nil {
			resp = Response{"20012","GetUnderCorrectedPaperByUserIdAndTestId  fail",err}
			c.Data["json"] = resp
			return
		}
		record.Score = sum
		record.Test_id = testId
		record.Test_record_type = underTest.Test_question_type
		record.User_id = userId
		record.Question_id=underTest.Question_id

		err = record.Save()
		if err!=nil {
			resp = Response{"20013","Save  fail",err}
			c.Data["json"] = resp
			return
		}
		err = underTest.SelfMarkDelete()
		if err!=nil {
			resp = Response{"20014","Delete  fail",err}
			c.Data["json"] = resp
			return
		}

		if math.Abs(float64(sum-test.Examiner_first_score))>float64(standardError) {
			var newUnderTest models.UnderCorrectedPaper
			newUnderTest.User_id="10000"
			newUnderTest.Self_score_id=userId
			newUnderTest.Test_id=testId
			newUnderTest.Question_id=test.Question_id
			newUnderTest.Test_question_type=7
			newUnderTest.Save()
		}

	}else if userId==test.Examiner_second_id {
		var sum int64
		//给试卷详情表打分
		for i := 0; i < len(testDetailIds); i++ {
			//取出小题试卷id,和小题分数
			var testInfo models.TestPaperInfo
			testDetailIdString:=testDetailIds[i]
			testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
			scoreString:=scores[i]
			score, _ := strconv.ParseInt(scoreString, 10, 64)
			//------------------------------------------------


			//查试卷小题
			err := testInfo.GetTestPaperInfo(testDetailId)
			if err != nil {
				resp := Response{"10008", "get testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			//修改试卷详情表


			testInfo.Examiner_second_self_score=score

			err = testInfo.Update()
			if err != nil {
				resp := Response{"10009", "update testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			sum += score
		}
		//给试卷表打分

		test.Examiner_second_self_score = sum

		err = test.Update()
		if err != nil {
			resp := Response{"10007", "update test fail", err}
			c.Data["json"] = resp
			return
		}
		//删除试卷待批改表 ，增加试卷记录表
		var record models.ScoreRecord
		var underTest models.UnderCorrectedPaper

		err = models.GetSelfScorePaperByTestQuestionTypeAndTestId(&underTest, testId,userId)
		if err!=nil {
			resp = Response{"20012","GetUnderCorrectedPaperByUserIdAndTestId  fail",err}
			c.Data["json"] = resp
			return
		}
		record.Score = sum
		record.Test_id = testId
		record.Test_record_type = underTest.Test_question_type
		record.User_id = userId
		record.Question_id=underTest.Question_id
		err = record.Save()
		if err!=nil {
			resp = Response{"20013","Save  fail",err}
			c.Data["json"] = resp
			return
		}
		err = underTest.SelfMarkDelete()
		if err!=nil {
			resp = Response{"20014","Delete  fail",err}
			c.Data["json"] = resp
			return
		}
		if math.Abs(float64(sum-test.Examiner_second_score))>float64(standardError) {
			var newUnderTest models.UnderCorrectedPaper
			newUnderTest.User_id="10000"
			newUnderTest.Test_id=testId
			newUnderTest.Self_score_id=userId
			newUnderTest.Question_id=test.Question_id
			newUnderTest.Test_question_type=7
			newUnderTest.Save()
		}

	} else if userId ==test.Examiner_third_id {
		var sum int64
		//给试卷详情表打分
		for i := 0; i < len(testDetailIds); i++ {
			//取出小题试卷id,和小题分数
			var testInfo models.TestPaperInfo
			testDetailIdString:=testDetailIds[i]
			testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
			scoreString:=scores[i]
			score, _ := strconv.ParseInt(scoreString, 10, 64)
			//------------------------------------------------


			//查试卷小题
			err := testInfo.GetTestPaperInfo(testDetailId)
			if err != nil {
				resp := Response{"10008", "get testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			//修改试卷详情表


			testInfo.Examiner_third_self_score=score

			err = testInfo.Update()
			if err != nil {
				resp := Response{"10009", "update testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			sum += score
		}
		//给试卷表打分

		test.Examiner_third_self_score = sum

		err = test.Update()
		if err != nil {
			resp := Response{"10007", "update test fail", err}
			c.Data["json"] = resp
			return
		}
		//删除试卷待批改表 ，增加试卷记录表
		var record models.ScoreRecord
		var underTest models.UnderCorrectedPaper

		err = models.GetSelfScorePaperByTestQuestionTypeAndTestId(&underTest, testId,userId)
		if err!=nil {
			resp = Response{"20012","GetUnderCorrectedPaperByUserIdAndTestId  fail",err}
			c.Data["json"] = resp
			return
		}
		record.Score = sum
		record.Test_id = testId
		record.Test_record_type = underTest.Test_question_type
		record.User_id = userId
		record.Question_id=underTest.Question_id

		err = record.Save()
		if err!=nil {
			resp = Response{"20013","Save  fail",err}
			c.Data["json"] = resp
			return
		}
		err = underTest.SelfMarkDelete()
		if err!=nil {
			resp = Response{"20014","Delete  fail",err}
			c.Data["json"] = resp
			return
		}
		if math.Abs(float64(sum-test.Examiner_third_score))>float64(standardError) {
			var newUnderTest models.UnderCorrectedPaper
			newUnderTest.User_id="10000"
			newUnderTest.Test_id=testId
			newUnderTest.Question_id=test.Question_id
			newUnderTest.Self_score_id=userId
			newUnderTest.Test_question_type=7
			newUnderTest.Save()
		}

	}


	//----------------------------------------
	resp = Response{"10000", "OK", nil}
	c.Data["json"] = resp
}
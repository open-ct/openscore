package controllers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/open-ct/openscore/model"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *ApiController) Display() {
	defer c.ServeJSON()
	var requestBody TestRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.TestId

	var testPaper model.TestPaper
	var topic model.Topic
	var subTopic []model.SubTopic
	var response TestDisplayResponse

	_, err = testPaper.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	err = topic.GetTopic(testPaper.QuestionId)
	if err != nil {
		resp := Response{"10003", "get topic fail", err}
		c.Data["json"] = resp
		return
	}
	err = model.GetSubTopicsByTestId(testPaper.QuestionId, &subTopic)
	if err != nil {
		resp := Response{"10004", "get subtopic fail", err}
		c.Data["json"] = resp
		return
	}

	for i := 0; i < len(subTopic); i++ {
		var testPaperInfo model.TestPaperInfo
		err = testPaperInfo.GetTestPaperInfoByTestIdAndQuestionDetailId(testId, subTopic[i].QuestionDetailId)
		if err != nil {
			resp := Response{"10005", "get testPaperInfo fail", err}
			c.Data["json"] = resp
			return
		}
		tempSubTopic := SubTopicPlus{SubTopic: subTopic[i], TestDetailId: testPaperInfo.TestDetailId}

		response.SubTopics = append(response.SubTopics, tempSubTopic)
		picName := testPaperInfo.PicSrc
		// 图片地址拼接 ，按服务器
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
		tempTestPaperInfo := TestPaperInfoPlus{TestPaperInfo: testPaperInfo, PicCode: encoding}
		response.TestInfos = append(response.TestInfos, tempTestPaperInfo)
	}
	response.QuestionId = topic.QuestionId
	response.QuestionName = topic.QuestionName
	response.TestId = testId
	resp := Response{"10000", "OK", response}
	c.Data["json"] = resp
}

func (c *ApiController) List() {
	defer c.ServeJSON()
	var response TestListResponse

	userId, err := c.GetSessionUserId()

	if err != nil {
		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
		c.Data["json"] = resp
		return
	}

	err = model.GetDistributedTestIdPaperByUserId(userId, &response.TestId)
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

	resp := Response{"10000", "OK", response}
	c.Data["json"] = resp

}

func (c *ApiController) Point() {
	defer c.ServeJSON()
	var requestBody TestPoint
	var resp Response
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}

	userId, err := c.GetSessionUserId()
	if err != nil {
		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
		c.Data["json"] = resp
		return
	}
	scoresStr := requestBody.Scores
	testId := requestBody.TestId
	testDetailIdStr := requestBody.TestDetailId
	scores := strings.Split(scoresStr, "-")
	testDetailIds := strings.Split(testDetailIdStr, "-")
	// -------------------------------------------------------

	// 获取该试卷大题 和抽象大题信息
	var test model.TestPaper
	var topic model.Topic
	_, err = test.GetTestPaper(testId)
	if err != nil || test.TestId == 0 {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	err = topic.GetTopic(test.QuestionId)
	if err != nil || topic.QuestionId == 0 {
		resp := Response{"10003", "get topic fail", err}
		c.Data["json"] = resp
		return
	}
	// 获取试卷未批改表信息（试卷批改状态类型）
	var underTest model.UnderCorrectedPaper
	err = underTest.GetUnderCorrectedPaper(userId, testId)
	if err != nil || underTest.QuestionId == 0 {
		resp := Response{"10004", "get underCorrected fail", err}
		c.Data["json"] = resp
		return
	}
	if underTest.TestQuestionType == 0 {
		standardError := topic.StandardError

		// 分三种情况
		if userId == test.ExaminerFirstId {
			var sum int64
			// 给试卷详情表打分
			for i := 0; i < len(testDetailIds); i++ {
				// 取出小题试卷id,和小题分数
				var testInfo model.TestPaperInfo
				testDetailIdString := testDetailIds[i]
				testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
				scoreString := scores[i]
				score, _ := strconv.ParseInt(scoreString, 10, 64)
				// ------------------------------------------------

				// 查试卷小题
				err := testInfo.GetTestPaperInfo(testDetailId)
				if err != nil {
					resp := Response{"10008", "get testPaper fail", err}
					c.Data["json"] = resp
					return
				}
				// 修改试卷详情表

				testInfo.ExaminerFirstSelfScore = score

				err = testInfo.Update()
				if err != nil {
					resp := Response{"10009", "update testPaper fail", err}
					c.Data["json"] = resp
					return
				}
				sum += score
			}
			// 给试卷表打分

			test.ExaminerFirstSelfScore = sum
			err = test.Update()
			if err != nil {
				resp := Response{"10007", "update test fail", err}
				c.Data["json"] = resp
				return
			}

			// 删除试卷待批改表 ，增加试卷记录表
			var record model.ScoreRecord
			var underTest model.UnderCorrectedPaper

			err = model.GetSelfScorePaperByTestQuestionTypeAndTestId(&underTest, testId, userId)
			if err != nil {
				resp = Response{"20012", "GetUnderCorrectedPaperByUserIdAndTestId  fail", err}
				c.Data["json"] = resp
				return
			}
			record.Score = sum
			record.TestId = testId
			record.TestRecordType = underTest.TestQuestionType
			record.UserId = userId
			record.QuestionId = underTest.QuestionId

			err = record.Save()
			if err != nil {
				resp = Response{"20013", "Save  fail", err}
				c.Data["json"] = resp
				return
			}
			err = underTest.SelfMarkDelete()
			if err != nil {
				resp = Response{"20014", "Delete  fail", err}
				c.Data["json"] = resp
				return
			}

			if math.Abs(float64(sum-test.ExaminerFirstScore)) > float64(standardError) {
				var newUnderTest model.UnderCorrectedPaper
				newUnderTest.UserId = 10000
				newUnderTest.SelfScoreId = userId
				newUnderTest.TestId = testId
				newUnderTest.QuestionId = test.QuestionId
				newUnderTest.TestQuestionType = 7
				newUnderTest.Save()
			}

		} else if userId == test.ExaminerSecondId {
			var sum int64
			// 给试卷详情表打分
			for i := 0; i < len(testDetailIds); i++ {
				// 取出小题试卷id,和小题分数
				var testInfo model.TestPaperInfo
				testDetailIdString := testDetailIds[i]
				testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
				scoreString := scores[i]
				score, _ := strconv.ParseInt(scoreString, 10, 64)
				// ------------------------------------------------

				// 查试卷小题
				err := testInfo.GetTestPaperInfo(testDetailId)
				if err != nil {
					resp := Response{"10008", "get testPaper fail", err}
					c.Data["json"] = resp
					return
				}
				// 修改试卷详情表

				testInfo.ExaminerSecondSelfScore = score

				err = testInfo.Update()
				if err != nil {
					resp := Response{"10009", "update testPaper fail", err}
					c.Data["json"] = resp
					return
				}
				sum += score
			}
			// 给试卷表打分

			test.ExaminerSecondSelfScore = sum

			err = test.Update()
			if err != nil {
				resp := Response{"10007", "update test fail", err}
				c.Data["json"] = resp
				return
			}
			// 删除试卷待批改表 ，增加试卷记录表
			var record model.ScoreRecord
			var underTest model.UnderCorrectedPaper

			err = model.GetSelfScorePaperByTestQuestionTypeAndTestId(&underTest, testId, userId)
			if err != nil {
				resp = Response{"20012", "GetUnderCorrectedPaperByUserIdAndTestId  fail", err}
				c.Data["json"] = resp
				return
			}
			record.Score = sum
			record.TestId = testId
			record.TestRecordType = underTest.TestQuestionType
			record.UserId = userId
			record.QuestionId = underTest.QuestionId
			err = record.Save()
			if err != nil {
				resp = Response{"20013", "Save  fail", err}
				c.Data["json"] = resp
				return
			}
			err = underTest.SelfMarkDelete()
			if err != nil {
				resp = Response{"20014", "Delete  fail", err}
				c.Data["json"] = resp
				return
			}
			if math.Abs(float64(sum-test.ExaminerSecondScore)) > float64(standardError) {
				var newUnderTest model.UnderCorrectedPaper
				newUnderTest.UserId = 10000
				newUnderTest.TestId = testId
				newUnderTest.SelfScoreId = userId
				newUnderTest.QuestionId = test.QuestionId
				newUnderTest.TestQuestionType = 7
				newUnderTest.Save()
			}

		} else if userId == test.ExaminerThirdId {
			var sum int64
			// 给试卷详情表打分
			for i := 0; i < len(testDetailIds); i++ {
				// 取出小题试卷id,和小题分数
				var testInfo model.TestPaperInfo
				testDetailIdString := testDetailIds[i]
				testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
				scoreString := scores[i]
				score, _ := strconv.ParseInt(scoreString, 10, 64)
				// ------------------------------------------------

				// 查试卷小题
				err := testInfo.GetTestPaperInfo(testDetailId)
				if err != nil {
					resp := Response{"10008", "get testPaper fail", err}
					c.Data["json"] = resp
					return
				}
				// 修改试卷详情表

				testInfo.ExaminerThirdSelfScore = score

				err = testInfo.Update()
				if err != nil {
					resp := Response{"10009", "update testPaper fail", err}
					c.Data["json"] = resp
					return
				}
				sum += score
			}
			// 给试卷表打分

			test.ExaminerThirdSelfScore = sum

			err = test.Update()
			if err != nil {
				resp := Response{"10007", "update test fail", err}
				c.Data["json"] = resp
				return
			}
			// 删除试卷待批改表 ，增加试卷记录表
			var record model.ScoreRecord
			var underTest model.UnderCorrectedPaper

			err = model.GetSelfScorePaperByTestQuestionTypeAndTestId(&underTest, testId, userId)
			if err != nil {
				resp = Response{"20012", "GetUnderCorrectedPaperByUserIdAndTestId  fail", err}
				c.Data["json"] = resp
				return
			}
			record.Score = sum
			record.TestId = testId
			record.TestRecordType = underTest.TestQuestionType
			record.UserId = userId
			record.QuestionId = underTest.QuestionId

			err = record.Save()
			if err != nil {
				resp = Response{"20013", "Save  fail", err}
				c.Data["json"] = resp
				return
			}
			err = underTest.SelfMarkDelete()
			if err != nil {
				resp = Response{"20014", "Delete  fail", err}
				c.Data["json"] = resp
				return
			}
			if math.Abs(float64(sum-test.ExaminerThirdScore)) > float64(standardError) {
				var newUnderTest model.UnderCorrectedPaper
				newUnderTest.UserId = 10000
				newUnderTest.TestId = testId
				newUnderTest.QuestionId = test.QuestionId
				newUnderTest.SelfScoreId = userId
				newUnderTest.TestQuestionType = 7
				newUnderTest.Save()
			}

		}

	} else { // score数组string转int
		var scoreArr []int64
		var sum int64 = 0
		var record model.ScoreRecord
		for _, i := range scores {
			j, err := strconv.ParseInt(i, 10, 64)
			sum += j
			if err != nil {
				panic(err)
			}
			scoreArr = append(scoreArr, j)
		}

		final := false

		if topic.ScoreType == 1 {
			test.ExaminerFirstId = userId
			test.ExaminerFirstScore = sum
			final = true
		} else if underTest.TestQuestionType == 2 && test.ExaminerFirstId == -1 {
			test.ExaminerFirstId = userId
			test.ExaminerFirstScore = sum
		} else if underTest.TestQuestionType == 2 && test.ExaminerSecondId == -1 {
			test.ExaminerSecondId = userId
			test.ExaminerSecondScore = sum
			if math.Abs(float64(test.ExaminerSecondScore)-float64(test.ExaminerFirstScore)) <= float64(topic.StandardError) {
				log.Println(math.Abs(float64(test.ExaminerSecondScore) - float64(test.ExaminerFirstScore)))
				sum = int64(math.Abs(float64(test.ExaminerSecondScore+test.ExaminerFirstScore)) / 2)
				final = true
			} else {
				newUnderTest := model.UnderCorrectedPaper{}
				// 随机 抽一个 人

				newUnderTest.UserId = model.FindNewUserId(test.ExaminerFirstId, test.ExaminerSecondId, test.QuestionId)
				newUnderTest.TestQuestionType = 3
				newUnderTest.TestId = underTest.TestId
				newUnderTest.QuestionId = underTest.QuestionId
				err = newUnderTest.Save()
				if err != nil {
					resp := Response{"10005", "insert undertest fail", err}
					c.Data["json"] = resp
					return
				}
			}
		}
		if underTest.TestQuestionType == 0 {

			test.LeaderId = userId
			test.LeaderScore = sum
			final = true
		}
		if underTest.TestQuestionType == 3 {
			test.ExaminerThirdId = userId
			test.ExaminerThirdScore = sum
			first := math.Abs(float64(test.ExaminerThirdScore - test.ExaminerFirstScore))
			second := math.Abs(float64(test.ExaminerThirdScore - test.ExaminerSecondScore))
			var small float64
			if first <= second {
				small = first
				sum = (test.ExaminerThirdScore + test.ExaminerFirstScore) / 2
			} else {
				small = second
				sum = (test.ExaminerThirdScore + test.ExaminerSecondScore) / 2
			}
			if small <= float64(topic.StandardError) {
				final = true
			} else {

				test.QuestionStatus = 2

				newUnderTest := model.UnderCorrectedPaper{}
				// 阅卷组长类型默认 id 10000
				newUnderTest.UserId = 10000
				newUnderTest.TestQuestionType = 4
				newUnderTest.TestId = underTest.TestId
				newUnderTest.QuestionId = underTest.QuestionId
				err = newUnderTest.Save()
				if err != nil {
					resp := Response{"10006", "insert undertest fail", err}
					c.Data["json"] = resp
					return
				}
			}
		}
		if final {
			test.FinalScore = sum
			record.TestFinish = 1
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
			var tempTest model.TestPaperInfo
			id, _ := strconv.ParseInt(testDetailIds[i], 10, 64)
			log.Println(id)
			err = tempTest.GetTestPaperInfo(id)
			if err != nil {
				resp := Response{"10008", "get testPaper fail", err}
				c.Data["json"] = resp
				return
			}
			if topic.ScoreType == 1 {
				tempTest.ExaminerFirstId = userId
				tempTest.ExaminerFirstScore = score
			} else if topic.ScoreType == 2 && tempTest.ExaminerFirstId == -1 {
				tempTest.ExaminerFirstId = userId
				tempTest.ExaminerFirstScore = score
			} else if topic.ScoreType == 2 && tempTest.ExaminerSecondId == -1 {
				tempTest.ExaminerSecondId = userId
				tempTest.ExaminerSecondScore = score
			}
			if underTest.TestQuestionType == 4 || underTest.TestQuestionType == 5 {
				tempTest.LeaderId = userId
				tempTest.LeaderScore = score
			} else if underTest.TestQuestionType == 3 {
				tempTest.ExaminerThirdId = userId
				tempTest.ExaminerThirdScore = score
			}
			if final {
				tempTest.FinalScore = score

			}
			err = tempTest.Update()
			if err != nil {
				resp := Response{"10009", "update testPaper fail", err}
				c.Data["json"] = resp
				return
			}
		}

		record.Score = sum
		record.QuestionId = topic.QuestionId
		record.TestId = testId
		record.TestRecordType = underTest.TestQuestionType
		record.UserId = userId
		record.ScoreTime = time.Now()
		err = record.Save()
		if err != nil {
			resp := Response{"10010", "insert record fail", err}
			c.Data["json"] = resp
			return
		}

	}

	resp = Response{"10000", "ok", err}
	c.Data["json"] = resp
}

func (c *ApiController) Problem() {
	defer c.ServeJSON()
	var requestBody TestProblem
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}

	userId, err := c.GetSessionUserId()
	if err != nil {
		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
		c.Data["json"] = resp
		return
	}
	problemType := requestBody.ProblemType
	testId := requestBody.TestId
	problemMessage := requestBody.ProblemMessage

	var underTest model.UnderCorrectedPaper
	var record model.ScoreRecord
	var test model.TestPaper

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

	newUnderTest.UserId = userId
	newUnderTest.TestQuestionType = 6
	newUnderTest.ProblemType = problemType
	newUnderTest.ProblemMessage = problemMessage
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

		test.QuestionStatus = 3
		err = test.Update()
		if err != nil {
			resp := Response{"10005", "update testPaper fail", err}
			c.Data["json"] = resp
			return
		}
	}

	record.TestRecordType = 5
	record.TestId = testId
	record.UserId = userId
	record.QuestionId = test.QuestionId
	record.TestRecordType = 5
	record.ProblemType = problemType
	err = record.Save()
	if err != nil {
		resp := Response{"10006", "insert record fail", err}
		c.Data["json"] = resp
		return
	}
	resp := Response{"10000", "ok", err}
	c.Data["json"] = resp
}

func (c *ApiController) Answer() {
	defer c.ServeJSON()
	var requestBody TestRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.TestId
	var test model.TestPaper
	_, err = test.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	var answerTest model.TestPaper
	err = answerTest.GetTestPaperByQuestionIdAndQuestionStatus(test.QuestionId, 5)
	if err != nil {
		resp := Response{"10003", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}

	var as TestAnswerResponse
	var tempString []string
	err = model.GetTestInfoPicListByTestId(answerTest.TestId, &tempString)
	if err != nil {
		resp := Response{"10004", "get testPaperInfo fail", err}
		c.Data["json"] = resp
		return
	}
	// 改成base64编码
	for i := 0; i < len(tempString); i++ {
		picName := tempString[i]
		// 图片地址拼接 ，按服务器
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

func (c *ApiController) ExampleDetail() {
	c.GetSessionUserId()
	defer c.ServeJSON()
	var requestBody ExampleDetail
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.ExampleTestId
	log.Println(testId)
	// ____________________________________________________________
	var test model.TestPaper
	_, err = test.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	var exampleTest []model.TestPaper
	// ??
	err = model.GetTestPaperListByQuestionIdAndQuestionStatus(test.QuestionId, 6, &exampleTest)
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

	var topic model.Topic
	err = topic.GetTopic(exampleTest[0].QuestionId)
	if err != nil {
		resp := Response{"10005", "get topic fail", err}
		c.Data["json"] = resp
		return
	}
	var response ExampleDetailResponse
	response.QuestionName = topic.QuestionName
	for i := 0; i < len(exampleTest); i++ {
		var temp []model.TestPaperInfo
		err = model.GetTestInfoListByTestId(exampleTest[i].TestId, &temp)
		if err != nil {
			resp := Response{"10006", "get testPaperInfo fail", err}
			c.Data["json"] = resp
			return
		}
		// 转64编码
		for j := 0; j < len(temp); j++ {
			picName := temp[j].PicSrc
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
			temp[j].PicSrc = encoding
		}
		response.Test = append(response.Test, temp)
	}
	resp := Response{"10000", "ok", response}
	c.Data["json"] = resp

}

func (c *ApiController) ExampleList() {
	defer c.ServeJSON()
	var requestBody TestRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	testId := requestBody.TestId
	// ----------------------------------------------------------------------
	var testPaper model.TestPaper
	_, err = testPaper.GetTestPaper(testId)
	if err != nil {
		resp := Response{"10002", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	var response ExampleListResponse
	err = model.GetTestPaperListByQuestionIdAndQuestionStatus(testPaper.QuestionId, 6, &response.TestPapers)
	if err != nil {
		resp := Response{"10003", "get testPaper fail", err}
		c.Data["json"] = resp
		return
	}
	resp := Response{"10000", "ok", response}
	c.Data["json"] = resp

}

func (c *ApiController) Review() {
	defer c.ServeJSON()
	var response TestReviewResponse

	userId, err := c.GetSessionUserId()
	if err != nil {
		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
		c.Data["json"] = resp
		return
	}

	var records []model.ScoreRecord
	err = model.GetLatestRecords(userId, &records)
	if err != nil {
		resp := Response{"10002", "get record fail", err}
		c.Data["json"] = resp
		return
	}
	for i := 0; i < len(records); i++ {
		response.TestId = append(response.TestId, records[i].TestId)
		response.Score = append(response.Score, records[i].Score)
		response.ScoreTime = append(response.ScoreTime, records[i].ScoreTime)
	}
	resp := Response{"10000", "ok", response}
	c.Data["json"] = resp
}

func (c *ApiController) ReviewPoint() {
	defer c.ServeJSON()
	var requestBody TestPoint
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp := Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}

	userId, err := c.GetSessionUserId()
	if err != nil {
		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
		c.Data["json"] = resp
		return
	}
	scoresstr := requestBody.Scores
	testId := requestBody.TestId
	testDetailIdstr := requestBody.TestDetailId
	scores := strings.Split(scoresstr, "-")
	testDetailIds := strings.Split(testDetailIdstr, "-")
	var scoreArr []int64
	var sum int64 = 0
	var record model.ScoreRecord
	for _, i := range scores {
		j, err := strconv.ParseInt(i, 10, 64)
		sum += j
		if err != nil {
			panic(err)
		}
		scoreArr = append(scoreArr, j)
	}

	var test model.TestPaper
	_, err = test.GetTestPaper(testId)
	if err != nil || test.TestId == 0 {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	// 判断是否二次阅卷
	var topic model.Topic
	topic.GetTopic(test.QuestionId)
	scoreType := topic.ScoreType

	num := 0
	if test.ExaminerFirstId == userId {
		num = 0
		test.ExaminerFirstScore = sum
		if scoreType == 1 {
			test.FinalScore = sum
			record.TestFinish = 1
		}
		var record model.ScoreRecord
		record.GetRecordByTestId(testId, userId)
		record.Score = sum
		record.Update()

	} else if test.ExaminerSecondId == userId {
		num = 1
		test.ExaminerSecondScore = sum
		var record model.ScoreRecord
		record.GetRecordByTestId(testId, userId)
		record.Score = sum
		record.Update()
	} else {
		num = 2
		test.ExaminerThirdScore = sum
		var record model.ScoreRecord
		record.GetRecordByTestId(testId, userId)
		record.Score = sum
		record.Update()
	}
	err = test.Update()
	if err != nil || test.TestId == 0 {
		resp := Response{"10003", "update test paper fail", err}
		c.Data["json"] = resp
		return
	}

	for i := 0; i < len(testDetailIds); i++ {
		var testInfo model.TestPaperInfo
		testInfoId, _ := strconv.ParseInt(testDetailIds[i], 10, 64)
		testInfo.GetTestPaperInfo(testInfoId)
		if num == 0 {
			testInfo.ExaminerFirstScore = scoreArr[i]
			if scoreType == 1 {
				testInfo.FinalScore = scoreArr[i]
			}
		} else if num == 1 {
			testInfo.ExaminerSecondScore = scoreArr[i]
		} else {
			testInfo.ExaminerThirdScore = scoreArr[i]
		}
		err = testInfo.Update()
		if err != nil || test.TestId == 0 {
			resp := Response{"10004", "update testinfo paper fail", err}
			c.Data["json"] = resp
			return
		}
	}
	c.Data["json"] = Response{"10000", "ok", nil}
}

// 自评列表 chen
func (c *ApiController) SelfScoreList() {
	defer c.ServeJSON()
	var response TestListResponse

	userId, err := c.GetSessionUserId()
	if err != nil {
		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
		c.Data["json"] = resp
		return
	}

	err = model.GetUnMarkSelfTestIdPaperByUserId(userId, &response.TestId)
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

// /**
// 20.自评卷打分
// */
// func (c *ApiController) SelfMarkPoint() {
// 	defer c.ServeJSON()
// 	var requestBody TestPoint
//
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
// 	if err != nil {
// 		c.Data["json"] = Response{"10001", "cannot unmarshal", err}
// 		return
// 	}
// 	userId, err := c.GetSessionUserId()
// 	if err != nil {
// 		resp := Response{Status: "10001", Msg: "get user info fail", Data: err}
// 		c.Data["json"] = resp
// 		return
// 	}
// 	testId := requestBody.TestId
// 	scoreStr := requestBody.Scores
// 	testDetailIdStr := requestBody.TestDetailId
// 	testDetailIds := strings.Split(testDetailIdStr, "-")
// 	scores := strings.Split(scoreStr, "-")
//
// 	// ---------------------------------------------------------------------------------------
//
// 	// 查找大题
// 	var test model.TestPaper
// 	_, err = test.GetTestPaper(testId)
// 	if err != nil || test.TestId == 0 {
// 		resp := Response{"10002", "get test paper fail", err}
// 		c.Data["json"] = resp
// 		return
// 	}
// 	var topic model.Topic
// 	topic.GetTopic(test.QuestionId)
//
// 	// ----------------------------------------
// 	c.Data["json"] = Response{"10000", "OK", nil}
// }

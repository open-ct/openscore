package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"openscore/model"
	"openscore/util"
	"os"
	"strconv"
	"strings"
	"time"
)

// 用户增删改查 TODO

// ReadUserExcel 导入用户
// func (c *ApiController) ReadUserExcel() {
// 	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
// 	defer c.ServeJSON()
//
// 	file, header, err := c.GetFile("excel")
//
// 	if err != nil {
// 		log.Println(err)
// 		c.Data["json"] = Response{Status: "10001", Msg: "cannot unmarshal", Data: err}
// 		return
// 	}
// 	tempFile, err := os.Create(header.Filename)
// 	io.Copy(tempFile, file)
// 	f, err := excelize.OpenFile(header.Filename)
// 	if err != nil {
// 		log.Println(err)
// 		c.Data["json"] = Response{Status: "30000", Msg: "excel 表导入错误", Data: err}
// 		return
// 	}
//
// 	// Get all the rows in the Sheet1.
// 	rows, err := f.GetRows("Sheet1")
// 	if err != nil {
// 		log.Println(err)
// 		c.Data["json"] = Response{Status: "30000", Msg: "excel 表导入错误", Data: err}
// 		return
// 	}
//
// 	for _, r := range rows[1:] {
// 		row := make([]string, len(rows[0]))
// 		copy(row, r)
// 		var user model.User
// 		user.UserName = row[0]
// 		user.IdCard = row[1]
// 		user.Password = row[2]
// 		user.Tel = row[3]
// 		user.Address = row[4]
// 		user.SubjectName = row[5]
// 		user.Email = row[6]
// 		userType, _ := strconv.Atoi(row[7])
// 		user.UserType = int64(userType)
// 		if err := user.Insert(); err != nil {
// 			log.Println(err)
// 			c.Data["json"] = Response{Status: "30001", Msg: "用户导入错误", Data: err}
// 			return
// 		}
//
// 	}
//
// 	err = tempFile.Close()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = os.Remove(header.Filename)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	// ------------------------------------------------
// 	data := make(map[string]interface{})
// 	data["data"] = nil
// 	c.Data["json"] = Response{Status: "10000", Msg: "OK", Data: data}
// }

/**
2.试卷导入
*/
func (c *ApiController) ReadExcel() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	defer c.ServeJSON()
	var resp Response

	file, header, err := c.GetFile("excel")

	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	tempFile, err := os.Create(header.Filename)
	io.Copy(tempFile, file)
	f, err := excelize.OpenFile(header.Filename)
	if err != nil {
		log.Println(err)
		resp = Response{"30000", "excel 表导入错误", err}
		c.Data["json"] = resp
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Println(err)
		resp = Response{"30000", "excel 表导入错误", err}
		c.Data["json"] = resp
		return
	}

	var bigQuestions []question
	var smallQuestions []question

	// 处理第一行 获取大题和小题的分布情况
	for i := 8; i < len(rows[0]); i++ {
		questionIds := strings.Split(rows[0][i], "-")

		id0, _ := strconv.Atoi(questionIds[0])
		id1, _ := strconv.Atoi(questionIds[1])
		if len(bigQuestions) > 0 && bigQuestions[len(bigQuestions)-1].Id == id0 {
			bigQuestions[len(bigQuestions)-1].Num++
		} else {
			bigQuestions = append(bigQuestions, question{Id: id0, Num: 1})
		}

		if len(smallQuestions) > 0 && smallQuestions[len(smallQuestions)-1].FatherId == id0 && smallQuestions[len(smallQuestions)-1].Id == id1 {
			smallQuestions[len(smallQuestions)-1].Num++
		} else {
			smallQuestions = append(smallQuestions, question{Id: id1, FatherId: id0, Num: 1})
		}
	}

	for _, smallQuestion := range smallQuestions {
		var topic model.Topic
		topic.QuestionId = int64(smallQuestion.Id)
		topic.ImportNumber = int64(len(rows) - 1)

		if err := topic.Update(); err != nil {
			log.Println(err)
			resp = Response{"30003", "大题导入试卷数更新错误", err}
			c.Data["json"] = resp
			return
		}
	}

	for _, r := range rows[1:] {
		row := make([]string, len(rows[0]))
		copy(row, r)
		index := 0
		smallIndex := 0
		// 处理该行的大题
		for _, bigQuestion := range bigQuestions {
			var testPaper model.TestPaper
			testPaper.TicketId = row[0]
			testPaper.QuestionId = int64(bigQuestion.Id)
			testPaper.Mobile = row[1]
			isParent, _ := strconv.Atoi(row[2])
			testPaper.IsParent = int64(isParent)
			testPaper.ClientIp = row[3]
			testPaper.Tag = row[4]
			testPaper.Candidate = row[6]
			testPaper.School = row[7]

			testId, err := testPaper.Insert()
			if err != nil {
				log.Println(err)
				resp = Response{"30001", "试卷大题导入错误", err}
				return
			}
			// 处理该大题的小题
			for num := smallIndex + bigQuestion.Num; smallIndex < num; smallIndex++ {
				content := row[index+8]
				for n := index + smallQuestions[smallIndex].Num - 1; index < n; index++ {

					content += "\n" + row[index+9]
					num--
				}

				src := util.UploadPic(row[0]+rows[0][8+index], content)

				var testPaperInfo model.TestPaperInfo
				testPaperInfo.TicketId = row[0]
				testPaperInfo.PicSrc = src

				testPaperInfo.TestId = testId
				testPaperInfo.QuestionDetailId = int64(smallQuestions[smallIndex].Id)

				if err := testPaperInfo.Insert(); err != nil {
					log.Println(err)
					resp = Response{"30002", "试卷小题导错误", err}
					return
				}
				index++
			}

		}
	}

	err = tempFile.Close()
	if err != nil {
		log.Println(err)
	}
	err = os.Remove(header.Filename)
	if err != nil {
		log.Println(err)
	}

	// ------------------------------------------------
	resp = Response{"10000", "OK", nil}
	c.Data["json"] = resp

}

type question struct {
	Id       int
	Num      int
	FatherId int
}

/**
样卷导入
*/

func (c *ApiController) ReadExampleExcel() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	defer c.ServeJSON()
	var resp Response
	var err error

	// ----------------------------------------------------

	file, header, err := c.GetFile("excel")
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	tempFile, err := os.Create(header.Filename)
	io.Copy(tempFile, file)
	f, err := excelize.OpenFile(header.Filename)
	if err != nil {
		log.Println(err)
		resp = Response{"30000", "excel 表导入错误", err}
		c.Data["json"] = resp
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet2")
	if err != nil {
		log.Println(err)
		resp = Response{"30000", "excel 表导入错误", err}
		c.Data["json"] = resp
		return
	}

	for i := 1; i < len(rows); i++ {
		for j := 1; j < len(rows[i]); j++ {

			if i >= 1 && j >= 3 {
				// 准备数据
				testIdStr := rows[i][0]
				testId, _ := strconv.ParseInt(testIdStr, 10, 64)
				questionIds := strings.Split(rows[0][j], "-")
				questionIdStr := questionIds[0]
				questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)
				questionDetailIdStr := questionIds[3]
				questionDetailId, _ := strconv.ParseInt(questionDetailIdStr, 10, 64)
				name := rows[i][2]
				// 填充数据
				var testPaperInfo model.TestPaperInfo
				var testPaper model.TestPaper

				testPaperInfo.QuestionDetailId = questionDetailId
				s := rows[i][j]
				// split := strings.Split(s, "\n")
				src := util.UploadPic(rows[i][0]+rows[0][j], s)
				testPaperInfo.PicSrc = src
				// 查看大题试卷是否已经导入
				has, err := testPaper.GetTestPaper(testId)
				if err != nil {
					log.Println(err)
				}

				// 导入大题试卷
				if !has {
					testPaper.TestId = testId
					testPaper.QuestionId = questionId
					testPaper.QuestionStatus = 6
					testPaper.Candidate = name
					_, err = testPaper.Insert()
					if err != nil {
						log.Println(err)
						resp = Response{"30001", "试卷大题导入错误", err}
						c.Data["json"] = resp
						return
					}
				}
				// 导入小题试卷
				testPaperInfo.TestId = testId
				err = testPaperInfo.Insert()
				if err != nil {
					log.Println(err)
					resp = Response{"30002", "试卷小题导错误", err}
					c.Data["json"] = resp
					return
				}

			}

		}

	}
	// 获取选项名 存导入试卷数
	for k := 3; k < len(rows[0]); k++ {
		questionIds := strings.Split(rows[0][k], "-")
		questionIdStr := questionIds[0]
		questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)
		var topic model.Topic
		topic.QuestionId = questionId
		topic.ImportNumber = int64(len(rows) - 1)
		err = topic.Update()
		if err != nil {
			log.Println(err)
			resp = Response{"30003", "大题导入试卷数更新错误", err}
			c.Data["json"] = resp
			return
		}
	}

	err = tempFile.Close()
	if err != nil {
		log.Println(err)
	}
	err = os.Remove(header.Filename)
	if err != nil {
		log.Println(err)
	}

	// ------------------------------------------------
	data := make(map[string]interface{})
	data["data"] = nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

func (c *ApiController) ReadAnswerExcel() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	defer c.ServeJSON()
	var resp Response
	var err error

	// ----------------------------------------------------

	file, header, err := c.GetFile("excel")
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	tempFile, err := os.Create(header.Filename)
	io.Copy(tempFile, file)
	f, err := excelize.OpenFile(header.Filename)
	if err != nil {
		log.Println(err)
		resp = Response{"30000", "excel 表导入错误", err}
		c.Data["json"] = resp
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet2")
	if err != nil {
		log.Println(err)
		resp = Response{"30000", "excel 表导入错误", err}
		c.Data["json"] = resp
		return
	}

	for i := 1; i < len(rows); i++ {
		for j := 1; j < len(rows[i]); j++ {

			if i >= 1 && j >= 3 {
				// 准备数据
				testIdStr := rows[i][0]
				testId, _ := strconv.ParseInt(testIdStr, 10, 64)
				questionIds := strings.Split(rows[0][j], "-")
				questionIdStr := questionIds[0]
				questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)
				questionDetailIdStr := questionIds[3]
				questionDetailId, _ := strconv.ParseInt(questionDetailIdStr, 10, 64)
				name := rows[i][2]
				// 填充数据
				var testPaperInfo model.TestPaperInfo
				var testPaper model.TestPaper

				testPaperInfo.QuestionDetailId = questionDetailId
				s := rows[i][j]
				// split := strings.Split(s, "\n")
				src := util.UploadPic(rows[i][0]+rows[0][j], s)
				testPaperInfo.PicSrc = src
				// 查看大题试卷是否已经导入
				has, err := testPaper.GetTestPaper(testId)
				if err != nil {
					log.Println(err)
				}

				// 导入大题试卷
				if !has {
					testPaper.TestId = testId
					testPaper.QuestionId = questionId
					testPaper.QuestionStatus = 5
					testPaper.Candidate = name
					_, err = testPaper.Insert()
					if err != nil {
						log.Println(err)
						resp = Response{"30001", "试卷大题导入错误", err}
						c.Data["json"] = resp
						return
					}
				}
				// 导入小题试卷
				testPaperInfo.TestId = testId
				err = testPaperInfo.Insert()
				if err != nil {
					log.Println(err)
					resp = Response{"30002", "试卷小题导错误", err}
					c.Data["json"] = resp
					return
				}

			}

		}

	}
	// 获取选项名 存导入试卷数
	for k := 3; k < len(rows[0]); k++ {
		questionIds := strings.Split(rows[0][k], "-")
		questionIdStr := questionIds[0]
		questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)
		var topic model.Topic
		topic.QuestionId = questionId
		topic.ImportNumber = int64(len(rows) - 1)
		err = topic.Update()
		if err != nil {
			log.Println(err)
			resp = Response{"30003", "大题导入试卷数更新错误", err}
			c.Data["json"] = resp
			return
		}
	}

	err = tempFile.Close()
	if err != nil {
		log.Println(err)
	}
	err = os.Remove(header.Filename)
	if err != nil {
		log.Println(err)
	}

	// ------------------------------------------------
	data := make(map[string]interface{})
	data["data"] = nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
3.大题列表
*/

func (c *ApiController) QuestionBySubList() {
	defer c.ServeJSON()
	var requestBody QuestionBySubList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	subjectName := requestBody.SubjectName
	// ----------------------------------------------------
	// 获取大题列表
	topics := make([]model.Topic, 0)
	err = model.FindTopicBySubNameList(&topics, subjectName)
	if err != nil {
		log.Println(err)
		resp = Response{"30004", "获取大题列表错误  ", err}
		c.Data["json"] = resp
		return
	}

	var questions = make([]QuestionBySubListVO, len(topics))
	for i := 0; i < len(topics); i++ {

		questions[i].QuestionId = topics[i].QuestionId
		questions[i].QuestionName = topics[i].QuestionName

	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["questionsList"] = questions
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
4.试卷参数导入
*/

func (c *ApiController) InsertTopic() {

	defer c.ServeJSON()
	var requestBody AddTopic
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// adminId := requestBody.AdminId
	topicName := requestBody.TopicName
	scoreType := requestBody.ScoreType
	score := requestBody.Score
	standardError := requestBody.Error
	subjectName := requestBody.SubjectName
	details := requestBody.TopicDetails

	// ----------------------------------------------------
	// 添加subject
	var subject model.Subject
	subject.SubjectName = subjectName
	flag, err := model.GetSubjectBySubjectName(&subject, subjectName)
	if err != nil {
		log.Println(err)
	}
	subjectId := subject.SubjectId
	if !flag {
		err, subjectId = model.InsertSubject(&subject)
		if err != nil {
			log.Println(err)
			resp = Response{"30005", "科目导入错误  ", err}
			c.Data["json"] = resp
			return
		}
	}
	// 添加topic
	var topic model.Topic
	topic.QuestionName = topicName
	topic.ScoreType = scoreType
	topic.QuestionScore = score
	topic.StandardError = standardError
	topic.SubjectName = subjectName
	topic.ImportTime = time.Now()
	topic.SubjectId = subjectId

	err, questionId := model.InsertTopic(&topic)
	if err != nil {
		log.Println(err)
		resp = Response{"30006", " 大题参数导入错误 ", err}
		c.Data["json"] = resp
		return
	}

	var addTopicVO AddTopicVO
	var addTopicDetailVOList = make([]AddTopicDetailVO, len(details))

	for i := 0; i < len(details); i++ {
		var subTopic model.SubTopic
		subTopic.QuestionDetailName = details[i].TopicDetailName
		subTopic.QuestionDetailScore = details[i].DetailScore
		subTopic.ScoreType = details[i].DetailScoreTypes
		subTopic.QuestionId = questionId
		err, questionDetailId := model.InsertSubTopic(&subTopic)
		if err != nil {
			log.Println(err)
			resp = Response{"30007", "小题参数导入错误  ", err}
			c.Data["json"] = resp
			return
		}
		addTopicDetailVOList[i].QuestionDetailId = questionDetailId
	}
	addTopicVO.QuestionId = questionId
	addTopicVO.QuestionDetailIds = addTopicDetailVOList
	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["addTopicVO"] = addTopicVO
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
5.科目选择
*/

func (c *ApiController) SubjectList() {

	defer c.ServeJSON()
	var resp Response

	// supervisorId := requestBody.SupervisorId
	// ----------------------------------------------------
	// 获取科目列表
	subjects := make([]model.Subject, 0)
	err := model.FindSubjectList(&subjects)
	if err != nil {
		log.Println(err)
		resp = Response{"30008", "科目列表获取错误  ", err}
		c.Data["json"] = resp
		return
	}

	var subjectVOList = make([]SubjectListVO, len(subjects))
	for i := 0; i < len(subjects); i++ {

		subjectVOList[i].SubjectName = subjects[i].SubjectName
		subjectVOList[i].SubjectId = subjects[i].SubjectId
	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["subjectVOList"] = subjectVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
6.试卷分配界面
*/
func (c *ApiController) DistributionInfo() {

	defer c.ServeJSON()
	var requestBody DistributionInfo
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ----------------------------------------------------
	// 标注输出
	var distributionInfoVO DistributionInfoVO
	// 获取试卷导入数量
	var topic model.Topic
	topic.QuestionId = questionId
	err = topic.GetTopic(questionId)
	if err != nil {
		log.Println(err)
		resp = Response{"30009", "获取试卷导入数量错误  ", err}
		c.Data["json"] = resp
		return
	}

	scoreType := topic.ScoreType
	distributionInfoVO.ScoreType = scoreType

	importNumber := topic.ImportNumber
	distributionInfoVO.ImportTestNumber = importNumber
	// 获取试卷未分配数量
	// 查询相应试卷
	papers := make([]model.TestPaper, 0)
	if err := model.FindUnDistributeTest(questionId, &papers); err != nil {
		log.Println(err)
		resp = Response{"30012", "试卷分配异常，无法获取未分配试卷 ", err}
		c.Data["json"] = resp
		return
	}

	distributionInfoVO.LeftTestNumber = len(papers)
	// 获取在线人数

	// 查找在线且未分配试卷的人
	usersList := make([]model.User, 0)
	if err := model.FindUsers(&usersList, topic.SubjectName); err != nil {
		log.Println(err)
		resp = Response{"30010", "获取可分配人数错误  ", err}
		c.Data["json"] = resp
		return
	}
	distributionInfoVO.OnlineNumber = int64(len(usersList))

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["distributionInfoVO"] = distributionInfoVO
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
7.试卷分配
*/
func (c *ApiController) Distribution() {

	defer c.ServeJSON()
	var requestBody Distribution
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	testNumber := requestBody.TestNumber
	userNumber := requestBody.UserNumber
	// ----------------------------------------------------

	// 是否需要二次阅卷
	var topic model.Topic
	topic.QuestionId = questionId
	err = topic.GetTopic(questionId)
	if err != nil {
		log.Println(err)
		resp = Response{"30011", "试卷分配异常,无法获取试卷批改次数 ", err}
		c.Data["json"] = resp
		return
	}
	scoreType := topic.ScoreType

	// 查询相应试卷
	papers := make([]model.TestPaper, 0)
	err = model.FindUnDistributeTest(questionId, &papers)

	if err != nil {
		log.Println(err)
		resp = Response{"30012", "试卷分配异常，无法获取未分配试卷 ", err}
		c.Data["json"] = resp
		return
	}
	testPapers := papers[:testNumber]

	// 查找在线且未分配试卷的人
	usersList := make([]model.User, 0)
	err = model.FindUsers(&usersList, topic.SubjectName)
	if err != nil {
		log.Println(err)
		resp = Response{"30013", "试卷分配异常，无法获取可分配阅卷员 ", err}
		c.Data["json"] = resp
		return
	}
	users := usersList[:userNumber]

	// 第一次分配试卷
	countUser := make([]int, userNumber)
	var ii int
	for i := 0; i < len(testPapers); {
		ii = i
		for j := 0; j < len(users); j++ {

			// 修改testPaper改为已分配
			testPapers[ii].CorrectingStatus = 1
			err := testPapers[ii].Update()
			if err != nil {
				log.Println(err)
				resp = Response{"30014", "试卷第一次分配异常，无法更改试卷状态 ", err}
				c.Data["json"] = resp
				return
			}

			// 添加试卷未批改记录
			var underCorrectedPaper model.UnderCorrectedPaper
			underCorrectedPaper.TestId = testPapers[ii].TestId
			underCorrectedPaper.QuestionId = testPapers[ii].QuestionId
			underCorrectedPaper.TestQuestionType = 1
			underCorrectedPaper.UserId = users[j].UserId
			if err := underCorrectedPaper.Save(); err != nil {
				log.Println(err)
				resp = Response{"30015", "试卷第一次分配异常，无法生成待批改试卷 ", err}
				c.Data["json"] = resp
				return
			}

			countUser[j]++
			testNumber--
			ii++

		}
		i += userNumber
	}

	// 修改user变为已分配
	for _, user := range users {
		user.GetUser(user.UserId)
		fmt.Println("user: ", user)

		user.IsDistribute = true
		user.QuestionId = questionId
		if err := user.UpdateCols("is_distribute", "question_id"); err != nil {
			log.Println(err)
			resp = Response{"30019", "试卷分配异常，用户分配状态更新失败 ", err}
			c.Data["json"] = resp
			return
		}
	}

	// 二次阅卷
	if scoreType == 2 {
		testNumber = len(testPapers)
		revers(users)
		var ii int
		for i := 0; i < len(testPapers); {
			ii = i
			for j := 0; j < len(users); j++ {
				if testNumber == 0 {
					break
				} else {
					// 修改testPaper改为已分配
					testPapers[ii].CorrectingStatus = 1
					err := testPapers[ii].Update()
					if err != nil {
						log.Println(err)
						resp = Response{"30016", "试卷第二次分配异常，无法更改试卷状态 ", err}
						c.Data["json"] = resp
						return
					}

					// 添加试卷未批改记录
					var underCorrectedPaper model.UnderCorrectedPaper
					underCorrectedPaper.TestId = testPapers[ii].TestId
					underCorrectedPaper.QuestionId = testPapers[ii].QuestionId
					underCorrectedPaper.TestQuestionType = 2
					underCorrectedPaper.UserId = users[j].UserId
					err = underCorrectedPaper.Save()
					if err != nil {
						log.Println(err)
						resp = Response{"30017", "试卷第二次分配异常，无法更改试卷状态 ", err}
						c.Data["json"] = resp
						return
					}
					countUser[j] = countUser[j] + 1
					testNumber--
					ii++
				}
			}
			i += userNumber
		}

	}

	for i := 0; i < userNumber; i++ {
		// 添加试卷分配表
		var paperDistribution model.PaperDistribution
		paperDistribution.TestDistributionNumber = int64(countUser[i])
		paperDistribution.UserId = users[i].UserId
		paperDistribution.QuestionId = questionId
		err := paperDistribution.Save()
		if err != nil {
			log.Println(err)
			resp = Response{"30018", "试卷分配异常，试卷分配添加异常 ", err}
			c.Data["json"] = resp
			return
		}
	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["data"] = nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
8.图片显示
*/
func (c *ApiController) Pic() {
	defer c.ServeJSON()
	var requestBody ReadFile
	var resp Response
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	// 获取图片名
	picName := requestBody.PicName
	// 获取图片地址（win版）
	src := "C:\\Users\\chen\\go\\src\\open-ct\\img\\" + picName
	// linux版（）
	// var src := "/usr/workspace/src/open-ct/"+name

	// -------------------------------------
	bytes, err := os.ReadFile(src)
	encoding := base64.StdEncoding.EncodeToString(bytes)
	if err != nil {
		log.Println(err)
		resp = Response{"30020", "图片显示异常 ", err}
		c.Data["json"] = resp
		return
	}
	// c.Ctx.Output.Header("Content-Type", "image/jpeg")
	// c.Ctx.Output.Header("Content-Length", strconv.Itoa(len(data)))
	// c.Ctx.WriteString(string(data))
	// c.Ctx.ResponseWriter.WriteHeader(200)
	// ----------------------------
	data := make(map[string]interface{})
	data["encoding"] = encoding
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
数组转置函数
*/
func revers(users []model.User) {
	for i := 0; i < len(users)/2; i++ {
		var temp model.User
		temp = users[i]
		users[i] = users[len(users)-i-1]
		users[len(users)-i-1] = temp
	}
}

func cutUser(oldData []model.User, n int) (newData []model.User) {
	newData1 := make([]model.User, n)
	for i := 0; i < n; i++ {
		newData1[i] = oldData[i]
	}
	return newData1
}

/**
9.大题展示列表
*/

func (c *ApiController) TopicList() {
	defer c.ServeJSON()
	var resp Response
	// supervisorId := requestBody.SupervisorId

	// ----------------------------------------------------
	// 获取大题列表
	topics := make([]model.Topic, 0)
	err := model.FindTopicList(&topics)
	if err != nil {
		log.Println(err)
		resp = Response{"30021", "获取大题参数设置记录表失败  ", err}
		c.Data["json"] = resp
		return
	}

	var topicVOList = make([]TopicVO, len(topics))
	for i := 0; i < len(topics); i++ {

		topicVOList[i].SubjectName = topics[i].SubjectName
		topicVOList[i].TopicName = topics[i].QuestionName
		topicVOList[i].Score = topics[i].QuestionScore
		topicVOList[i].StandardError = topics[i].StandardError
		topicVOList[i].ScoreType = topics[i].ScoreType
		topicVOList[i].TopicId = topics[i].QuestionId
		topicVOList[i].ImportTime = topics[i].ImportTime

		subTopics := make([]model.SubTopic, 0)
		model.FindSubTopicsByQuestionId(topics[i].QuestionId, &subTopics)
		if err != nil {
			log.Println(err)
			resp = Response{"30022", "获取小题参数设置记录表失败  ", err}
			c.Data["json"] = resp
			return
		}
		subTopicVOS := make([]SubTopicVO, len(subTopics))
		for j := 0; j < len(subTopics); j++ {
			subTopicVOS[j].SubTopicId = subTopics[j].QuestionDetailId
			subTopicVOS[j].SubTopicName = subTopics[j].QuestionDetailName
			subTopicVOS[j].Score = subTopics[j].QuestionDetailScore
			subTopicVOS[j].ScoreDistribution = subTopics[j].ScoreType
		}
		topicVOList[i].SubTopicVOList = subTopicVOS
	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["topicVOList"] = topicVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

// DistributionRecord ...
func (c *ApiController) DistributionRecord() {
	defer c.ServeJSON()
	var requestBody DistributionRecord
	var resp Response

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	subjectName := requestBody.SubjectName
	// ----------------------------------------------------
	// 获取大题列表
	topics := make([]model.Topic, 0)
	err = model.FindTopicBySubNameList(&topics, subjectName)
	if err != nil {
		log.Println(err)
		resp = Response{"30023", "获取试卷分配记录表失败  ", err}
		c.Data["json"] = resp
		return
	}

	var distributionRecordList = make([]DistributionRecordVO, len(topics))
	for i := 0; i < len(topics); i++ {

		distributionRecordList[i].TopicId = topics[i].QuestionId
		distributionRecordList[i].TopicName = topics[i].QuestionName
		distributionRecordList[i].ImportNumber = topics[i].ImportNumber
		distributionTestNumber, err := model.CountTestDistributionNumberByQuestionId(topics[i].QuestionId)
		if err != nil {
			log.Println(err)
			resp = Response{"30024", "获取试卷分配记录表失败，统计试卷已分配数失败  ", err}
			c.Data["json"] = resp
			return
		}
		distributionUserNumber, err := model.CountUserDistributionNumberByQuestionId(topics[i].QuestionId)
		if err != nil {
			log.Println(err)
			resp = Response{"30025", "获取试卷分配记录表失败，统计用户已分配数失败  ", err}
			c.Data["json"] = resp
			return
		}
		distributionRecordList[i].DistributionUserNumber = distributionUserNumber
		distributionRecordList[i].DistributionTestNumber = distributionTestNumber

	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["distributionRecordList"] = distributionRecordList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
试卷删除
*/

func (c *ApiController) DeleteTest() {

	defer c.ServeJSON()
	var requestBody DeleteTest
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// adminId := requestBody.AdminId
	questionId := requestBody.QuestionId

	// ----------------------------------------------------
	count, err := model.CountUnScoreTestNumberByQuestionId(questionId)
	if count == 0 {
		model.DeleteAllTest(questionId)
		subTopics := make([]model.SubTopic, 0)
		model.FindSubTopicsByQuestionId(questionId, &subTopics)
		for j := 0; j < len(subTopics); j++ {
			subTopic := subTopics[j]
			testPaperInfos := make([]model.TestPaperInfo, 0)
			model.FindTestPaperInfoByQuestionDetailId(subTopic.QuestionDetailId, &testPaperInfos)
			for k := 0; k < len(testPaperInfos); k++ {
				picName := testPaperInfos[k].PicSrc
				src := "./img/" + picName
				os.Remove(src)
				testPaperInfos[k].Delete()
			}
		}

	} else {
		resp = Response{"30030", "试卷未批改完不能删除  ", err}
		c.Data["json"] = resp
		return
	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["data"] = nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

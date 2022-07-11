package supervisor

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"math"
	. "openscore/controllers"
	"openscore/model"
	"strconv"
	"strings"
	"time"
)

type SupervisorApiController struct {
	beego.Controller
}

/**
9.大题选择列表
*/
func (c *SupervisorApiController) QuestionList() {
	defer c.ServeJSON()
	var requestBody QuestionList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	supervisorId := requestBody.SupervisorId
	// ----------------------------------------------------
	// 获取大题列表
	var user model.User
	user.User_id = supervisorId
	user.GetUser(supervisorId)
	subjectName := user.Subject_name

	topics := make([]model.Topic, 0)
	err = model.FindTopicBySubNameList(&topics, subjectName)
	fmt.Println("topics: ", topics, subjectName)

	if err != nil {
		resp = Response{"20000", "GetTopicList err ", err}
		c.Data["json"] = resp
		return
	}

	var questions = make([]QuestionListVO, len(topics))
	for i := 0; i < len(topics); i++ {

		questions[i].QuestionId = topics[i].Question_id
		questions[i].QuestionName = topics[i].Question_name

	}

	// ----------------------------------------------------
	data := make(map[string]interface{})
	data["questionsList"] = questions
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
10.用户登入信息表
*/
func (c *SupervisorApiController) UserInfo() {
	defer c.ServeJSON()
	var requestBody UserInfo
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	supervisorId := requestBody.SupervisorId

	// ----------------------------------------------------
	user := model.User{User_id: supervisorId}
	err = user.GetUser(supervisorId)
	if err != nil {
		resp = Response{"20001", "获取用户信息失败", err}
		c.Data["json"] = resp
		return
	}
	var userInfoVO UserInfoVO
	userInfoVO.UserName = user.User_name
	userInfoVO.SubjectName = user.Subject_name

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["userInfo"] = userInfoVO
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
8.教师监控页面
*/
func (c *SupervisorApiController) TeacherMonitoring() {
	defer c.ServeJSON()
	var requestBody TeacherMonitoring
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ----------------------------------------------------
	paperDistributions := make([]model.PaperDistribution, 0)
	err = model.FindPaperDistributionByQuestionId(&paperDistributions, questionId)
	if err != nil {
		resp = Response{"20021", "试卷分配信息获取失败  ", err}
		c.Data["json"] = resp
		return
	}
	teacherMonitoringList := make([]TeacherMonitoringVO, len(paperDistributions))
	for i := 0; i < len(paperDistributions); i++ {
		// 教师id
		userId := paperDistributions[i].User_id
		teacherMonitoringList[i].UserId = userId
		// 分配试卷数量
		testDistributionNumber := paperDistributions[i].Test_distribution_number
		teacherMonitoringList[i].TestDistributionNumber = testDistributionNumber
		// 试卷失败数
		failCount, err1 := model.CountFailTestNumberByUserId(userId, questionId)
		if err1 != nil {
			resp = Response{"20024", "获取试卷批改失败数 错误 ", err}
			c.Data["json"] = resp
			return
		}
		teacherMonitoringList[i].TestProblemNumber = failCount
		failCountString := strconv.FormatInt(failCount, 10)
		failCountFloat, _ := strconv.ParseFloat(failCountString, 64)

		// 试卷剩余未批改数
		remainingTestNumber, err1 := model.CountRemainingTestNumberByUserId(questionId, userId)
		if err1 != nil {
			resp = Response{"20023", "无法获取试卷未批改数", err}
			c.Data["json"] = resp
			return
		}
		teacherMonitoringList[i].TestRemainingNumber = remainingTestNumber

		// 试卷完成数
		finishCount := paperDistributions[i].Test_distribution_number - failCount - remainingTestNumber
		teacherMonitoringList[i].TestSuccessNumber = (float64(finishCount))
		// 用户信息
		user := model.User{User_id: userId}
		err = user.GetUser(userId)
		if err != nil {
			resp = Response{"20001", "无法获取用户信息", err}
			c.Data["json"] = resp
			return
		}
		// 用户名
		teacherMonitoringList[i].UserName = user.User_name
		// 是否在线
		isOnline := user.UserType
		teacherMonitoringList[i].IsOnline = isOnline
		// 在线时间
		var onlineTime int64
		if isOnline == 1 {
			endingTime := time.Now().Unix()
			startTime := user.Login_time.Unix()
			tempTime := endingTime - startTime
			onlineTime = user.Online_time + (tempTime)
		} else {
			onlineTime = user.Online_time
		}
		// 平均速度  (秒/份)
		var markingSpeed float64 = 99999999
		s1 := strconv.FormatInt(onlineTime, 10)
		s, _ := strconv.ParseFloat(s1, 64)

		if finishCount != 0 {
			tempSpeed := s / float64(finishCount)
			markingSpeed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tempSpeed), 64)
		}
		teacherMonitoringList[i].MarkingSpeed = markingSpeed
		// 预计时间 (小时)

		predictTime := markingSpeed * float64(remainingTestNumber)
		predictTime = predictTime / 3600
		predictTime, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", predictTime), 64)
		teacherMonitoringList[i].PredictTime = predictTime

		// 平均分
		var averageScore float64 = 0
		if finishCount != 0 {
			sum, err1 := model.SumFinishScore(userId, questionId)
			if err1 != nil {
				resp = Response{"20025", "计算平均分失败", err}
				c.Data["json"] = resp
				return
			}
			averageScore = sum / float64(finishCount)
		}

		teacherMonitoringList[i].AverageScore = averageScore
		// 有效度
		var validity float64 = 0
		if (float64(finishCount) + failCountFloat) != 0 {
			validity = float64(finishCount) / (float64(finishCount) + failCountFloat)
			validity, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", validity), 64)
		}
		teacherMonitoringList[i].Validity = validity
		// //自评率
		// selfTestCount ,err1 := models.CountSelfScore(userId,questionId)
		// if err1!=nil {
		//	resp = Response{"20026","CountSelfScore  fail",err}
		//	c.Data["json"] = resp
		//	return
		// }
		// selfTestCountString:=strconv.FormatInt(selfTestCount,10)
		// selfTestCountFloat,_:=strconv.ParseFloat(selfTestCountString,64)
		//
		// var selfScoreRate float64=0
		// if finishCount!=0 {
		//	selfScoreRate= selfTestCountFloat/float64(finishCount)
		// }
		// teacherMonitoringList[i].EvaluationIndex=selfScoreRate

		// 标准差
		var add float64
		finishScoreList := make([]model.ScoreRecord, 0)
		model.FindFinishTestByUserId(&finishScoreList, userId, questionId)
		for j := 0; j < len(finishScoreList); j++ {
			scoreJ := finishScoreList[j].Score
			tempJ := math.Abs((float64(scoreJ)) - averageScore)
			add = add + math.Exp2(tempJ)

		}
		sqrt := math.Sqrt(add)
		sqrt, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", sqrt), 64)
		teacherMonitoringList[i].StandardDeviation = sqrt

	}

	// --------------------------------------------------

	data := make(map[string]interface{})

	data["teacherMonitoringList"] = teacherMonitoringList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
11.分数分布表
*/
func (c *SupervisorApiController) ScoreDistribution() {
	defer c.ServeJSON()
	var requestBody ScoreDistribution
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ----------------------------------------------------
	// 求大题满分
	topic := model.Topic{Question_id: questionId}
	err = topic.GetTopic(questionId)
	if err != nil {
		resp = Response{"20002", "could not find topic", err}
		c.Data["json"] = resp
		return
	}
	questionScore := topic.Question_score
	// 求该大题的已批改试卷表
	scoreRecordList := make([]model.ScoreRecord, 0)
	err = model.FindFinishScoreRecordListByQuestionId(&scoreRecordList, questionId)
	if err != nil {
		resp = Response{"20003", "FindFinishScoreRecordListByQuestionId err", err}
		c.Data["json"] = resp
		return

	}
	// 该题已批改试卷总数
	count := len(scoreRecordList)
	countString := strconv.FormatInt(int64(count), 10)
	countFloat, _ := strconv.ParseFloat(countString, 64)
	// 标准的输出数据
	scoreDistributionList := make([]ScoreDistributionVO, questionScore+1)
	// 统计分数
	var i int64 = 0
	for ; i <= questionScore; i++ {
		scoreDistributionList[i].Score = i
		number, err := model.CountTestByScore(questionId, i)
		if err != nil {
			resp = Response{"20004", "CountTestByScore err", err}
			c.Data["json"] = resp
			return
		}

		numberString := strconv.FormatInt(number, 10)
		numberFloat, _ := strconv.ParseFloat(numberString, 64)
		scoreDistribution := 0.00
		if countFloat != 0 {
			scoreDistribution = numberFloat / countFloat
			scoreDistribution, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", scoreDistribution), 64)
		}
		scoreDistributionList[i].Rate = scoreDistribution
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["scoreDistributionList"] = scoreDistributionList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
12.大题教师选择列表
*/
func (c *SupervisorApiController) TeachersByQuestion() {
	defer c.ServeJSON()
	var requestBody TeachersByQuestion
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ----------------------------------------------------
	// 根据大题求试卷分配表
	paperDistributions := make([]model.PaperDistribution, 0)
	err = model.FindPaperDistributionByQuestionId(&paperDistributions, questionId)
	if err != nil {
		resp = Response{"20005", "FindPaperDistributionByQuestionId err", err}
		c.Data["json"] = resp
		return
	}

	// 输出标准
	teacherVOList := make([]TeacherVO, len(paperDistributions))

	// 求教师名和转化输出
	for i := 0; i < len(paperDistributions); i++ {
		userId := paperDistributions[i].User_id
		user := model.User{User_id: userId}
		err := user.GetUser(userId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		userName := user.User_name
		teacherVOList[i].UserId = userId
		teacherVOList[i].UserName = userName
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["teacherVOList"] = teacherVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
13.自评监控表
*/
func (c *SupervisorApiController) SelfScore() {
	defer c.ServeJSON()
	var requestBody SelfScore
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	examinerId := requestBody.ExaminerId
	// ----------------------------------------------------

	// 根据userId找到自评卷

	selfScoreRecord := make([]model.ScoreRecord, 0)
	model.FindSelfScoreRecordByUserId(&selfScoreRecord, examinerId)
	if err != nil {
		resp = Response{"20006", "FindSelfScoreRecordByUserId err", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	selfScoreRecordVOList := make([]SelfScoreRecordVO, len(selfScoreRecord))

	// 求教师名和转化输出
	for i := 0; i < len(selfScoreRecord); i++ {
		testId := selfScoreRecord[i].Test_id
		var test model.TestPaper
		test.GetTestPaperByTestId(testId)

		// 求大题信息 标准误差
		var topic model.Topic
		topic.GetTopic(test.Question_id)
		standardError := topic.Standard_error
		var error float64
		if test.Examiner_first_id == examinerId {
			selfScoreRecordVOList[i].Score = test.Examiner_first_score
			selfScoreRecordVOList[i].SelfScore = test.Examiner_first_self_score
			selfScoreRecordVOList[i].StandardError = standardError

			error = math.Abs(float64(selfScoreRecordVOList[i].Score - selfScoreRecordVOList[i].SelfScore))
			if error <= float64(standardError) {
				selfScoreRecordVOList[i].IsQualified = 1
			} else {
				selfScoreRecordVOList[i].IsQualified = 0
			}
		} else if test.Examiner_second_id == examinerId {
			selfScoreRecordVOList[i].Score = test.Examiner_second_score
			selfScoreRecordVOList[i].SelfScore = test.Examiner_second_self_score
			selfScoreRecordVOList[i].StandardError = standardError
			error = math.Abs(float64(selfScoreRecordVOList[i].Score - selfScoreRecordVOList[i].SelfScore))
			if error <= float64(standardError) {
				selfScoreRecordVOList[i].IsQualified = 1
			} else {
				selfScoreRecordVOList[i].IsQualified = 0
			}
		} else if test.Examiner_third_id == examinerId {
			selfScoreRecordVOList[i].Score = test.Examiner_third_score
			selfScoreRecordVOList[i].SelfScore = test.Examiner_third_self_score
			selfScoreRecordVOList[i].StandardError = standardError
			error = math.Abs(float64(selfScoreRecordVOList[i].Score - selfScoreRecordVOList[i].SelfScore))
			if error <= float64(standardError) {
				selfScoreRecordVOList[i].IsQualified = 1
			} else {
				selfScoreRecordVOList[i].IsQualified = 0
			}
		}
		selfScoreRecordVOList[i].TestId = testId
		selfScoreRecordVOList[i].Error = error

	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["selfScoreRecordVOList"] = selfScoreRecordVOList

	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
14，平均分监控表
*/
func (c *SupervisorApiController) AverageScore() {
	defer c.ServeJSON()
	var requestBody AverageScore
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	// --------------------------------------------
	// 根据大题求试卷分配表
	paperDistributions := make([]model.PaperDistribution, 0)
	err = model.FindPaperDistributionByQuestionId(&paperDistributions, questionId)
	if err != nil {
		resp = Response{"20007", "FindPaperDistributionByQuestionId err", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	scoreAverageVOList := make([]ScoreAverageVO, len(paperDistributions))
	var sumAllTestScore = 0.0
	var count = 0.0
	// 求教师名和转化输出
	for i := 0; i < len(paperDistributions); i++ {
		// 求userId 和userName
		userId := paperDistributions[i].User_id
		user := model.User{User_id: userId}
		err := user.GetUser(userId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		userName := user.User_name
		scoreAverageVOList[i].UserId = userId
		scoreAverageVOList[i].UserName = userName

		scoreNumber, err := model.CountTestScoreNumberByUserId(userId, questionId)
		if err != nil {
			resp = Response{"20008", "CountFinishTestNumberByUserId fail", err}
			c.Data["json"] = resp
			return
		}
		finishCountString := strconv.FormatInt(scoreNumber, 10)
		finishCountFloat, _ := strconv.ParseFloat(finishCountString, 64)
		count = count + finishCountFloat
		var averageScore float64 = 0
		if finishCountFloat != 0 {
			sum, err := model.SumFinishScore(userId, questionId)
			if err != nil {
				resp = Response{"20009", "SumFinishScore fail", err}
				c.Data["json"] = resp
				return
			}
			averageScore = sum / finishCountFloat
			sumAllTestScore = sumAllTestScore + sum
		}
		averageScore, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", averageScore), 64)
		scoreAverageVOList[i].Average = averageScore

	}
	var topic = model.Topic{Question_id: questionId}
	err = topic.GetTopic(questionId)
	if err != nil {
		resp = Response{"20000", "GetTopicList err ", err}
		c.Data["json"] = resp
		return
	}
	var fullScore = topic.Question_score
	var questionAverageScore = 0.0
	if count != 0 {
		questionAverageScore = sumAllTestScore / count
	}
	questionAverageScore, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", questionAverageScore), 64)

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["scoreAverageVOList"] = scoreAverageVOList
	data["fullScore"] = fullScore
	data["questionAverageScore"] = questionAverageScore
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
17，问题卷表
*/
func (c *SupervisorApiController) ProblemTest() {
	defer c.ServeJSON()
	var requestBody ProblemTest
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	// ------------------------------------------------

	// 根据大题号找到问题卷
	problemUnderCorrectedPaper := make([]model.UnderCorrectedPaper, 0)
	model.FindProblemUnderCorrectedPaperByQuestionId(&problemUnderCorrectedPaper, questionId)
	if err != nil {
		resp = Response{"20010", "FindProblemUnderCorrectedPaperByQuestionId  fail", err}
		c.Data["json"] = resp
		return
	}

	// 问题卷的数量
	var count = len(problemUnderCorrectedPaper)

	// 输出标准
	ProblemUnderCorrectedPaperVOList := make([]ProblemUnderCorrectedPaperVO, count)

	// 求阅卷老师名和转化输出
	for i := 0; i < len(problemUnderCorrectedPaper); i++ {
		// 存testId
		ProblemUnderCorrectedPaperVOList[i].TestId = problemUnderCorrectedPaper[i].Test_id
		// 存userId  userName
		userId := problemUnderCorrectedPaper[i].User_id
		user := model.User{User_id: userId}
		err := user.GetUser(userId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		userName := user.User_name
		ProblemUnderCorrectedPaperVOList[i].ExaminerId = userId
		ProblemUnderCorrectedPaperVOList[i].ExaminerName = userName
		// 存问题类型
		ProblemUnderCorrectedPaperVOList[i].ProblemType = problemUnderCorrectedPaper[i].Problem_type
		ProblemUnderCorrectedPaperVOList[i].ProblemMes = problemUnderCorrectedPaper[i].Problem_message

	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["ProblemUnderCorrectedPaperVOList"] = ProblemUnderCorrectedPaperVOList
	data["count"] = count
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
18，仲裁卷表
*/
func (c *SupervisorApiController) ArbitramentTest() {
	defer c.ServeJSON()
	var requestBody ArbitramentTest
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	// ------------------------------------------------

	// 根据大题号找到仲裁卷
	arbitramentUnderCorrectedPaper := make([]model.UnderCorrectedPaper, 0)
	err = model.FindArbitramentUnderCorrectedPaperByQuestionId(&arbitramentUnderCorrectedPaper, questionId)
	if err != nil {
		resp = Response{"20011", "FindArbitramentUnderCorrectedPaperByQuestionId  fail", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	arbitramentTestVOList := make([]ArbitramentTestVO, len(arbitramentUnderCorrectedPaper))

	var count = len(arbitramentUnderCorrectedPaper)
	// 求阅卷老师名和转化输出
	for i := 0; i < len(arbitramentUnderCorrectedPaper); i++ {
		// 存testId
		var testId = arbitramentUnderCorrectedPaper[i].Test_id
		arbitramentTestVOList[i].TestId = testId

		// 查试卷
		var testPaper model.TestPaper
		testPaper.Test_id = testId

		testPaper.GetTestPaper(testId)
		// 查存试卷第一次评分人id
		var examinerFirstId = testPaper.Examiner_first_id
		arbitramentTestVOList[i].ExaminerFirstId = examinerFirstId
		// 查第一次评分人
		firstExaminer := model.User{User_id: examinerFirstId}
		err := firstExaminer.GetUser(examinerFirstId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		// 查第一次评分人姓名
		examinerFirstName := firstExaminer.User_name
		// 存试卷第一次评分人姓名和分数
		arbitramentTestVOList[i].ExaminerFirstName = examinerFirstName
		arbitramentTestVOList[i].ExaminerFirstScore = testPaper.Examiner_first_score

		// 查存试卷第二次评分人id
		var examinerSecondId = testPaper.Examiner_second_id
		arbitramentTestVOList[i].ExaminerSecondId = examinerSecondId
		// 查第二次试卷评分人
		secondExaminer := model.User{User_id: examinerSecondId}
		err = secondExaminer.GetUser(examinerSecondId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		// 查第二次评分人姓名
		secondExaminerName := secondExaminer.User_name
		// 存第二次评分人姓名和分数
		arbitramentTestVOList[i].ExaminerSecondName = secondExaminerName
		arbitramentTestVOList[i].ExaminerSecondScore = testPaper.Examiner_second_score

		// 查存试卷第三次评分人id
		var examinerThirdId = testPaper.Examiner_third_id
		arbitramentTestVOList[i].ExaminerThirdId = examinerThirdId
		// 查第二次试卷评分人
		thirdExaminer := model.User{User_id: examinerThirdId}
		err = thirdExaminer.GetUser(examinerThirdId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		// 查第三次评分人姓名
		thirdExaminerName := thirdExaminer.User_name
		// 存第三次评分人姓名和分数
		arbitramentTestVOList[i].ExaminerThirdName = thirdExaminerName
		arbitramentTestVOList[i].ExaminerThirdScore = testPaper.Examiner_third_score
		// 查存实际误差
		arbitramentTestVOList[i].PracticeError = testPaper.Pratice_error
		// 查存标准误差
		var topic model.Topic
		topic.GetTopic(questionId)
		arbitramentTestVOList[i].StandardError = topic.Standard_error

	}
	// 查存该题满分
	var topic = model.Topic{Question_id: questionId}
	topic.GetTopic(questionId)
	var fullScore = topic.Question_score

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["arbitramentTestVOList"] = arbitramentTestVOList
	data["count"] = count
	data["fullScore"] = fullScore
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
15.总体进度
*/
func (c *SupervisorApiController) ScoreProgress() {
	defer c.ServeJSON()
	var requestBody ScoreProgress
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	subject := requestBody.Subject

	// ----------------------------------------------------
	// 根据科目获取大题列表
	topics := make([]model.Topic, 0)
	err = model.FindTopicBySubNameList(&topics, subject)
	if err != nil {
		resp = Response{"20000", "GetTopicList err ", err}
		c.Data["json"] = resp
		return
	}
	// 确定输出标准
	scoreProgressVOList := make([]ScoreProgressVO, len(topics))

	for i := 0; i < len(topics); i++ {
		// 获取大题id
		questionId := topics[i].Question_id
		scoreProgressVOList[i].QuestionId = questionId
		// 获取大题名
		questionName := topics[i].Question_name
		scoreProgressVOList[i].QuestionName = questionName
		// 自评率
		scoreProgressVOList[i].SelfScoreRate = topics[i].SelfScoreRate

		// 获取 任务总量
		importNumber := topics[i].Import_number
		scoreProgressVOList[i].ImportNumber = importNumber

		// 出成绩量
		finishNumber, err := model.CountFinishScoreNumberByQuestionId(questionId)
		if err != nil {
			resp = Response{"20013", "CountFinishScoreNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}
		scoreProgressVOList[i].FinishNumber = finishNumber

		// 出成绩率
		finishNumberString := strconv.FormatInt(finishNumber, 10)
		finishNumberFloat, _ := strconv.ParseFloat(finishNumberString, 64)
		importNumberString := strconv.FormatInt(importNumber, 10)
		importNumberFloat, _ := strconv.ParseFloat(importNumberString, 64)
		var finishRate float64 = 0
		if importNumberFloat != 0 {
			finishRate = finishNumberFloat / importNumberFloat
			finishRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", finishRate), 64)
		}
		scoreProgressVOList[i].FinishRate = finishRate
		// 未出成绩量
		unfinishedNumberFloat := importNumberFloat - finishNumberFloat
		scoreProgressVOList[i].UnfinishedNumber = unfinishedNumberFloat

		// 未出成绩率
		var unfinishedRate float64 = 0
		if importNumberFloat != 0 {
			unfinishedRate = unfinishedNumberFloat / importNumberFloat
			unfinishedRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", unfinishedRate), 64)
		}
		scoreProgressVOList[i].UnfinishedRate = unfinishedRate
		// 是否全部完成
		var isAllFinished string
		if unfinishedNumberFloat != 0 {
			isAllFinished = "未完成"
		} else {
			isAllFinished = "完成"
		}
		scoreProgressVOList[i].IsAllFinished = isAllFinished
		// 在线人数
		var users = make([]model.User, 0)
		model.FindUserNumberByQuestionId(&users, questionId)
		usersNumber := len(users)
		scoreProgressVOList[i].DistributionUserNumber = int64(usersNumber)

		// -------------------------------------------------------------------------------
		// 平均速度
		var onlineUserNumber int64 = 0
		// 全部用户在线时间
		var totalUserOnlineTime int64 = 0
		// 当前在线用户在线时间
		var currentUserOnlineTime int64 = 0
		// 全部用户批改量
		var totalUserFinishNumber = 0
		// 当前在线用户批改试卷量
		var currentUserFinishNumber = 0
		// 全部用户批改试卷总分
		var totalUserScoreSum int64 = 0
		// 当前用户 批改试卷总分
		var currentUserScoreSum int64 = 0

		for i := 0; i < usersNumber; i++ {
			userId := users[i].User_id
			isOnline := users[i].Status
			userOnlineTime := users[i].Online_time
			userRecord := make([]model.ScoreRecord, 0)
			model.FindFinishTestByUserId(&userRecord, userId, questionId)
			tempFinishNumber := len(userRecord)
			var tempScore int64 = 0
			for j := 0; j < len(userRecord); j++ {
				tempScore = tempScore + userRecord[j].Score
			}
			if isOnline == 1 {
				// 计算时间
				endingTime := time.Now().Unix()
				startTime := users[i].Login_time.Unix()
				tempTime := endingTime - startTime
				userOnlineTime = userOnlineTime + (tempTime)

				currentUserOnlineTime = currentUserOnlineTime + userOnlineTime
				// 计算在线任务量和分数
				currentUserFinishNumber = currentUserFinishNumber + tempFinishNumber
				currentUserScoreSum = currentUserScoreSum + tempScore
				onlineUserNumber = onlineUserNumber + 1
			}
			// 计算时间
			totalUserOnlineTime = totalUserOnlineTime + userOnlineTime
			// 计算任务量和分数
			totalUserFinishNumber = totalUserFinishNumber + tempFinishNumber
			totalUserScoreSum = totalUserScoreSum + tempScore
		}
		// 平均分
		var averageScore = 0.0
		if totalUserFinishNumber != 0 {
			averageScore = float64(totalUserScoreSum) / float64(totalUserFinishNumber)
		}
		// 在线用户平均分
		scoreProgressVOList[i].AverageScore = averageScore
		var currentAverageScore = 0.0
		if currentUserFinishNumber != 0 {
			currentAverageScore = float64(currentUserScoreSum) / float64(currentUserFinishNumber)
		}
		scoreProgressVOList[i].OnlineAverageScore = currentAverageScore
		// 阅卷速度
		var scoreSpeed = 999999999.0
		if totalUserFinishNumber != 0 {
			scoreSpeed = float64(totalUserOnlineTime) / float64(totalUserFinishNumber)
		}
		scoreProgressVOList[i].AverageSpeed = scoreSpeed
		// 在线阅卷速度
		var onlineScoreSpeed = 99999999.0
		if currentUserFinishNumber != 0 {
			onlineScoreSpeed = float64(currentUserOnlineTime) / float64(currentUserFinishNumber)
		}
		scoreProgressVOList[i].OnlineAverageSpeed = onlineScoreSpeed
		//  在线人数 预估时间
		scoreProgressVOList[i].OnlinePredictTime = onlineScoreSpeed * unfinishedNumberFloat
		// 全部人数预估时间
		scoreProgressVOList[i].PredictTime = scoreSpeed * unfinishedNumberFloat
		//
		scoreProgressVOList[i].OnlineUserNumber = onlineUserNumber

		// --------------------------------------------------
		// 一次评卷完成数
		firstScoreNumber, err1 := model.CountFirstScoreNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20014", "CountFirstScoreNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}
		scoreProgressVOList[i].FirstFinishedNumber = firstScoreNumber
		// 一次评卷完成率
		firstScoreNumberString := strconv.FormatInt(firstScoreNumber, 10)
		firstScoreNumberFloat, _ := strconv.ParseFloat(firstScoreNumberString, 64)
		var firstScoreRate float64 = 0
		if importNumberFloat != 0 {
			firstScoreRate = firstScoreNumberFloat / importNumberFloat
			firstScoreRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", firstScoreRate), 64)
		}
		scoreProgressVOList[i].FirstFinishedRate = firstScoreRate
		// 未出第一次成绩量
		firstUnfinishedNumber := importNumberFloat - firstScoreNumberFloat
		scoreProgressVOList[i].FirstUnfinishedNumber = firstUnfinishedNumber
		// 第一次未出成绩率
		var firstUnfinishedRate float64 = 0
		if importNumberFloat != 0 {
			firstUnfinishedRate = firstUnfinishedNumber / importNumberFloat
			firstUnfinishedRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", firstUnfinishedRate), 64)
		}
		scoreProgressVOList[i].FirstUnfinishedRate = firstUnfinishedRate
		// 第一次阅卷是否全部完成
		var isFirstFinished string
		if firstUnfinishedNumber != 0 {
			isFirstFinished = "未完成"
		} else {
			isFirstFinished = "完成"
		}
		scoreProgressVOList[i].IsFirstFinished = isFirstFinished

		// -----------------------------------------

		// 二次评卷完成数
		secondScoreNumber, err1 := model.CountSecondScoreNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20015", "CountSecondScoreNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}
		scoreProgressVOList[i].SecondFinishedNumber = secondScoreNumber
		// 二次评卷完成率
		secondScoreNumberString := strconv.FormatInt(secondScoreNumber, 10)
		secondScoreNumberFloat, _ := strconv.ParseFloat(secondScoreNumberString, 64)

		var secondScoreRate float64 = 0
		if importNumberFloat != 0 {
			secondScoreRate = secondScoreNumberFloat / importNumberFloat
			secondScoreRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", secondScoreRate), 64)
		}
		scoreProgressVOList[i].SecondFinishedRate = secondScoreRate

		// 未出第二次成绩量
		secondUnfinishedNumber := importNumberFloat - firstScoreNumberFloat
		scoreProgressVOList[i].SecondUnfinishedNumber = secondUnfinishedNumber
		// 第二次未出成绩率
		var secondUnfinishedRate float64 = 0
		if importNumberFloat != 0 {
			secondUnfinishedRate = secondUnfinishedNumber / importNumberFloat
			secondUnfinishedRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", secondUnfinishedRate), 64)
		}
		scoreProgressVOList[i].SecondUnfinishedRate = secondUnfinishedRate
		// 第二次阅卷是否全部完成
		var isSecondFinished string
		if secondUnfinishedNumber != 0 {
			isSecondFinished = "未完成"
		} else {
			isSecondFinished = "完成"
		}
		scoreProgressVOList[i].IsSecondFinished = isSecondFinished

		// -----------------------------------------

		// 三次评卷完成数
		thirdScoreNumber, err1 := model.CountThirdScoreNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20016", "CountThirdScoreNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}
		scoreProgressVOList[i].ThirdFinishedNumber = thirdScoreNumber
		// 三次评卷完成率
		thirdScoreNumberString := strconv.FormatInt(thirdScoreNumber, 10)
		thirdScoreNumberFloat, _ := strconv.ParseFloat(thirdScoreNumberString, 64)
		var thirdScoreRate float64 = 0
		if importNumberFloat != 0 {
			thirdScoreRate = thirdScoreNumberFloat / importNumberFloat
			thirdScoreRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", thirdScoreRate), 64)
		}
		scoreProgressVOList[i].ThirdFinishedRate = thirdScoreRate

		// 未出第三次成绩量
		thirdUnfinishedNumber := importNumberFloat - thirdScoreNumberFloat
		scoreProgressVOList[i].ThirdUnfinishedNumber = thirdUnfinishedNumber
		// 未出成绩率
		var thirdUnfinishedRate float64 = 0
		if importNumberFloat != 0 {
			thirdUnfinishedRate = thirdUnfinishedNumber / importNumberFloat
			thirdUnfinishedRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", thirdUnfinishedRate), 64)
		}
		scoreProgressVOList[i].ThirdUnfinishedRate = thirdUnfinishedRate
		// 第三次阅卷是否全部完成
		var isThirdFinished string
		if thirdUnfinishedNumber != 0 {
			isThirdFinished = "未完成"
		} else {
			isThirdFinished = "完成"
		}
		scoreProgressVOList[i].IsThirdFinished = isThirdFinished

		// -----------------------------------------
		// 仲裁卷完成数
		arbitramentFinishNumber, err1 := model.CountArbitramentFinishNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20017", "CountArbitramentFinishNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}
		scoreProgressVOList[i].ArbitramentFinishedNumber = arbitramentFinishNumber
		// 仲裁卷未完成量
		arbitramentUnfinishedNumber, err1 := model.CountArbitramentUnFinishNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20018", "CountArbitramentUnFinishNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}

		scoreProgressVOList[i].ArbitramentUnfinishedNumber = arbitramentUnfinishedNumber
		// 仲裁卷产生量：
		arbitramentNumber := arbitramentFinishNumber + arbitramentUnfinishedNumber
		scoreProgressVOList[i].ArbitramentNumber = arbitramentNumber
		// 仲裁卷产生率
		arbitramentNumberString := strconv.FormatInt(arbitramentNumber, 10)
		arbitramentNumberFloat, _ := strconv.ParseFloat(arbitramentNumberString, 64)
		var arbitramentRate float64 = 0
		if importNumberFloat != 0 {
			arbitramentRate = arbitramentNumberFloat / importNumberFloat
			arbitramentRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", arbitramentRate), 64)
		}
		scoreProgressVOList[i].ArbitramentRate = arbitramentRate

		// 仲裁卷完成率
		arbitramentFinishNumberString := strconv.FormatInt(arbitramentFinishNumber, 10)
		arbitramentFinishNumberFloat, _ := strconv.ParseFloat(arbitramentFinishNumberString, 64)
		var arbitramentFinishRate float64 = 0
		if arbitramentNumberFloat != 0 {
			arbitramentFinishRate = arbitramentFinishNumberFloat / arbitramentNumberFloat
			arbitramentFinishRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", arbitramentFinishRate), 64)
		}
		scoreProgressVOList[i].ArbitramentFinishedRate = arbitramentFinishRate

		// 仲裁卷未完成率
		arbitramentUnFinishNumberString := strconv.FormatInt(arbitramentUnfinishedNumber, 10)
		arbitramentUnFinishNumberFloat, _ := strconv.ParseFloat(arbitramentUnFinishNumberString, 64)
		var arbitramentUnfinishedRate float64 = 0
		if arbitramentNumberFloat != 0 {
			arbitramentUnfinishedRate = arbitramentUnFinishNumberFloat / arbitramentNumberFloat
			arbitramentUnfinishedRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", arbitramentUnfinishedRate), 64)
		}
		scoreProgressVOList[i].ArbitramentUnfinishedRate = arbitramentUnfinishedRate
		// 仲裁卷是否全部完成
		var ArbitramentFinished string
		if arbitramentUnfinishedNumber != 0 {
			ArbitramentFinished = "未完成"
		} else {
			ArbitramentFinished = "完成"
		}
		scoreProgressVOList[i].IsArbitramentFinished = ArbitramentFinished

		// -----------------------------------------

		// 问题卷完成数
		problemFinishNumber, err1 := model.CountProblemFinishNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20019", "CountProblemFinishNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}
		scoreProgressVOList[i].ProblemFinishedNumber = problemFinishNumber
		// 问题卷未完成量
		problemUnfinishedNumber, err1 := model.CountProblemUnFinishNumberByQuestionId(questionId)
		if err1 != nil {
			resp = Response{"20020", "CountProblemUnFinishNumberByQuestionId  fail", err}
			c.Data["json"] = resp
			return
		}

		scoreProgressVOList[i].ProblemUnfinishedNumber = problemUnfinishedNumber

		// 问题卷产生量：
		problemNumber := problemFinishNumber + problemUnfinishedNumber
		scoreProgressVOList[i].ProblemNumber = problemNumber

		// 问题卷产生率
		problemNumberString := strconv.FormatInt(problemNumber, 10)
		problemNumberFloat, _ := strconv.ParseFloat(problemNumberString, 64)
		var problemRate float64 = 0
		if importNumberFloat != 0 {
			problemRate = problemNumberFloat / importNumberFloat
			problemRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", problemRate), 64)
		}
		scoreProgressVOList[i].ProblemRate = problemRate

		// 问题卷完成率
		problemFinishedNumberString := strconv.FormatInt(problemFinishNumber, 10)
		problemFinishNumberFloat, _ := strconv.ParseFloat(problemFinishedNumberString, 64)
		var problemFinishRate float64 = 0
		if problemNumberFloat != 0 {
			problemFinishRate = problemFinishNumberFloat / problemNumberFloat
			problemFinishRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", problemFinishRate), 64)
		}
		scoreProgressVOList[i].ProblemFinishedRate = problemFinishRate

		// 问题卷未完成率
		problemUnfinishedNumberrString := strconv.FormatInt(problemUnfinishedNumber, 10)
		problemUnfinishedNumberFloat, _ := strconv.ParseFloat(problemUnfinishedNumberrString, 64)
		var problemUnfinishedRate float64 = 0
		if problemNumberFloat != 0 {
			problemUnfinishedRate = problemUnfinishedNumberFloat / problemNumberFloat
			problemUnfinishedRate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", problemUnfinishedRate), 64)
		}
		scoreProgressVOList[i].ProblemUnfinishedRate = problemUnfinishedRate
		// 问题卷是否全部完成
		var IsProblemFinished string
		if problemUnfinishedNumber != 0 {
			IsProblemFinished = "未完成"
		} else {
			IsProblemFinished = "完成"
		}
		scoreProgressVOList[i].IsProblemFinished = IsProblemFinished
	}
	// --------------------------------------------------

	data := make(map[string]interface{})

	data["scoreProgressVOList"] = scoreProgressVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
19.阅卷组长批改试卷
*/
func (c *SupervisorApiController) SupervisorPoint() {
	defer c.ServeJSON()
	var requestBody SupervisorPoint
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	supervisorId := requestBody.SupervisorId
	testId := requestBody.TestId
	scoreStr := requestBody.Scores
	testDetailIdStr := requestBody.TestDetailIds
	testDetailIds := strings.Split(testDetailIdStr, "-")
	scores := strings.Split(scoreStr, "-")

	// ---------------------------------------------------------------------------------------
	// 创建试卷小题详情

	var test model.TestPaper

	var sum int64
	// 给试卷详情表打分
	for i := 0; i < len(testDetailIds); i++ {
		// 取出小题试卷id,和小题分数
		var testInfo model.TestPaperInfo
		testDetailIdString := testDetailIds[i]
		testDetailId, _ := strconv.ParseInt(testDetailIdString, 10, 64)
		scoreString := scores[i]
		score, _ := strconv.ParseInt(scoreString, 10, 64)
		// 查试卷小题
		err := testInfo.GetTestPaperInfo(testDetailId)
		if err != nil {
			resp := Response{"10008", "get testPaper fail", err}
			c.Data["json"] = resp
			return
		}
		// 修改试卷详情表
		testInfo.Leader_id = supervisorId
		testInfo.Leader_score = score
		testInfo.Final_score = score
		testInfo.Final_score_id = supervisorId
		err = testInfo.Update()
		if err != nil {
			resp := Response{"10009", "update testPaper fail", err}
			c.Data["json"] = resp
			return
		}
		sum += score
	}
	// 给试卷表打分
	_, err = test.GetTestPaper(testId)
	if err != nil || test.Test_id == 0 {
		resp := Response{"10002", "get test paper fail", err}
		c.Data["json"] = resp
		return
	}
	test.Leader_id = supervisorId
	test.Leader_score = sum
	test.Final_score = sum
	test.Final_score_id = supervisorId
	err = test.Update()
	if err != nil {
		resp := Response{"10007", "update test fail", err}
		c.Data["json"] = resp
		return
	}
	// 删除试卷待批改表 ，增加试卷记录表
	var record model.ScoreRecord
	var underTest model.UnderCorrectedPaper

	err = model.GetUnderCorrectedSupervisorPaperByTestQuestionTypeAndTestId(&underTest, testId)
	if err != nil {
		resp = Response{"20012", "GetUnderCorrectedPaperByUserIdAndTestId  fail", err}
		c.Data["json"] = resp
		return
	}
	record.Score = sum
	record.Test_id = testId
	record.Test_record_type = underTest.Test_question_type
	record.User_id = supervisorId
	record.Question_id = underTest.Question_id
	record.Problem_type = underTest.Problem_type
	if underTest.Test_question_type != 7 {
		record.Test_finish = 1
	}

	err = record.Save()
	if err != nil {
		resp = Response{"20013", "Save  fail", err}
		c.Data["json"] = resp
		return
	}
	err = underTest.SupervisorDelete()
	if err != nil {
		resp = Response{"20014", "Delete  fail", err}
		c.Data["json"] = resp
		return
	}
	// ----------------------------------------
	resp = Response{"10000", "OK", nil}
	c.Data["json"] = resp
}

/**
20.问题卷列表
*/
func (c *SupervisorApiController) ProblemUnmarkList() {
	defer c.ServeJSON()
	var requestBody ProblemUnmarkList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	// ------------------------------------------------

	// 根据大题号找到问题卷
	problemUnderCorrectedPaper := make([]model.UnderCorrectedPaper, 0)
	model.FindProblemUnderCorrectedPaperByQuestionId(&problemUnderCorrectedPaper, questionId)
	if err != nil {
		resp = Response{"20027", "FindProblemUnderCorrectedList  fail", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	ProblemUnmarkVOList := make([]ProblemUnmarkListVO, len(problemUnderCorrectedPaper))

	// 求阅卷输出
	for i := 0; i < len(problemUnderCorrectedPaper); i++ {
		// 存testId
		ProblemUnmarkVOList[i].TestId = problemUnderCorrectedPaper[i].Test_id
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["ProblemUnmarkVOList"] = ProblemUnmarkVOList

	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
20.自评卷列表
*/
func (c *SupervisorApiController) SelfUnmarkList() {
	defer c.ServeJSON()
	var requestBody SelfUnmarkList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	// ------------------------------------------------

	// 根据大题号找到自评卷
	selfUnderCorrectedPaper := make([]model.UnderCorrectedPaper, 0)
	model.FindSelfUnderCorrectedPaperByQuestionId(&selfUnderCorrectedPaper, questionId)
	if err != nil {
		resp = Response{"20027", "FindSelfUnderCorrectedPaperByQuestionId  fail", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	selfUnmarkVOList := make([]SelfUnmarkListVO, len(selfUnderCorrectedPaper))

	// 求阅卷输出
	for i := 0; i < len(selfUnderCorrectedPaper); i++ {
		// 存testId
		selfUnmarkVOList[i].TestId = selfUnderCorrectedPaper[i].Test_id
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["selfUnmarkVOList"] = selfUnmarkVOList

	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

/**
21.仲裁卷列表
*/
func (c *SupervisorApiController) ArbitramentUnmarkList() {
	defer c.ServeJSON()
	var requestBody ArbitramentUnmarkList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ------------------------------------------------

	// 找到仲裁卷
	arbitramentUnderCorrectedPaper := make([]model.UnderCorrectedPaper, 0)
	err = model.FindAllArbitramentUnderCorrectedPaper(&arbitramentUnderCorrectedPaper, questionId)
	if err != nil {
		resp = Response{"20026", "FindAllArbitramentUnderCorrectedPaper  fail", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	arbitramentUnmarkListVOList := make([]ArbitramentUnmarkListVO, len(arbitramentUnderCorrectedPaper))

	for i := 0; i < len(arbitramentUnderCorrectedPaper); i++ {
		// 存testId
		var testId = arbitramentUnderCorrectedPaper[i].Test_id
		arbitramentUnmarkListVOList[i].TestId = testId
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["arbitramentUnmarkListVOList"] = arbitramentUnmarkListVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

// 16标准差

func (c *SupervisorApiController) ScoreDeviation() {
	defer c.ServeJSON()
	var requestBody ScoreDeviation
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ------------------------------------------------

	// 根据大题求试卷分配表
	paperDistributions := make([]model.PaperDistribution, 0)
	err = model.FindPaperDistributionByQuestionId(&paperDistributions, questionId)
	if err != nil {
		resp = Response{"20007", "FindPaperDistributionByQuestionId err", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	ScoreDeviationVOList := make([]ScoreDeviationVO, len(paperDistributions))
	userScoreNumbers := make([]int, len(paperDistributions))
	var count = 0

	// 求教师名和转化输出
	for i := 0; i < len(paperDistributions); i++ {
		// 求userId 和userName
		userId := paperDistributions[i].User_id
		user := model.User{User_id: userId}
		err := user.GetUser(userId)
		if err != nil {
			resp = Response{"20001", "could not found user", err}
			c.Data["json"] = resp
			return
		}
		userName := user.User_name
		ScoreDeviationVOList[i].UserId = userId
		ScoreDeviationVOList[i].UserName = userName

		var finishScoreList []model.ScoreRecord
		err = model.FindFinishTestByUserId(&finishScoreList, userId, questionId)
		if err != nil {
			resp = Response{"2027", "FindFinishTestNumberByUserId fail", err}
			c.Data["json"] = resp
			return
		}
		var finishCount = len(finishScoreList)
		userScoreNumbers[i] = finishCount
		count = count + finishCount

		var averageScore float64 = 0
		if finishCount != 0 {
			sum, err := model.SumFinishScore(userId, questionId)
			if err != nil {
				resp = Response{"20009", "SumFinishScore fail", err}
				c.Data["json"] = resp
				return
			}
			averageScore = math.Abs(sum / (float64(finishCount)))
		}

		var add float64
		for j := 0; j < finishCount; j++ {
			scoreJ := finishScoreList[j].Score
			tempJ := math.Abs((float64(scoreJ)) - averageScore)
			add = add + math.Exp2(tempJ)

		}
		sqrt := math.Sqrt(add)
		sqrt, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", sqrt), 64)
		ScoreDeviationVOList[i].DeviationScore = sqrt
	}
	QuestionScoreDeviation := 0.0
	for j := 0; j < len(ScoreDeviationVOList); j++ {
		QuestionScoreDeviation = QuestionScoreDeviation + (ScoreDeviationVOList[j].DeviationScore * (float64(userScoreNumbers[j]) / float64(count)))
		QuestionScoreDeviation, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", QuestionScoreDeviation), 64)
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["ScoreDeviationVOList"] = ScoreDeviationVOList
	data["QuestionScoreDeviation"] = QuestionScoreDeviation
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
22.自评卷列表
*/
func (c *SupervisorApiController) SelfMarkList() {
	defer c.ServeJSON()
	var requestBody SelfMarkList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}
	// supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	// ------------------------------------------------
	// 找到大题标准误差
	var topic model.Topic
	topic.GetTopic(questionId)
	standardError := topic.Standard_error
	// 找到自评卷
	selfMarkPaper := make([]model.UnderCorrectedPaper, 0)
	err = model.FindSelfMarkPaperByQuestionId(&selfMarkPaper, questionId)
	if err != nil {
		resp = Response{"20026", "FindAllArbitramentUnderCorrectedPaper  fail", err}
		c.Data["json"] = resp
		return
	}
	// 输出标准
	selfMarkVOList := make([]SelfMarkListVO, len(selfMarkPaper))

	for i := 0; i < len(selfMarkPaper); i++ {
		// 存testId
		testId := selfMarkPaper[i].Test_id
		selfScoreId := selfMarkPaper[i].Self_score_id

		var test model.TestPaper
		test.GetTestPaperByTestId(testId)

		if test.Examiner_first_id == selfScoreId {
			selfMarkVOList[i].Score = test.Examiner_first_score
			selfMarkVOList[i].SelfScore = test.Examiner_first_self_score

		} else if test.Examiner_second_id == selfScoreId {
			selfMarkVOList[i].Score = test.Examiner_second_score
			selfMarkVOList[i].SelfScore = test.Examiner_second_self_score
		} else if test.Examiner_third_id == selfScoreId {
			selfMarkVOList[i].Score = test.Examiner_third_score
			selfMarkVOList[i].SelfScore = test.Examiner_third_self_score
		}
		selfMarkVOList[i].Userid = selfScoreId
		var user model.User
		user.GetUser(selfScoreId)

		selfMarkVOList[i].Name = user.User_name
		selfMarkVOList[i].TestId = testId
		selfMarkVOList[i].StandardError = standardError
		selfMarkVOList[i].Error = math.Abs(float64(selfMarkVOList[i].Score - selfMarkVOList[i].SelfScore))
	}

	// --------------------------------------------------

	data := make(map[string]interface{})
	data["selfMarkVOList"] = selfMarkVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}

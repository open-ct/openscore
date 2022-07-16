package supervisor

import (
	"github.com/xuri/excelize/v2"
	"log"
	. "openscore/controllers"
	"openscore/model"
	"strconv"
)

// 导出成绩
func (c *SupervisorApiController) WriteScoreExcel() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	defer c.ServeJSON()

	f := excelize.NewFile()

	subjects := make([]model.Subject, 0)
	if err := model.FindSubjectList(&subjects); err != nil {
		log.Println(err)
		c.Data["json"] = Response{Status: "30008", Msg: "科目列表获取错误  ", Data: err}
		return
	}

	for _, subject := range subjects {
		topics := make([]model.Topic, 0)
		if err := model.FindTopicBySubNameList(&topics, subject.SubjectName); err != nil {
			log.Println(err)
			c.Data["json"] = Response{Status: "30004", Msg: "获取大题列表错误  ", Data: err}
			return
		}

		topicMap := make(map[int64]string, len(topics))
		// 获取subject的小题数
		var subjectSubTopics []model.SubTopic
		var subjectTestPapers []model.TestPaper
		for _, topic := range topics {
			topicMap[topic.QuestionId] = topic.QuestionName
			subTopics := make([]model.SubTopic, 0)
			if err := model.FindSubTopicsByQuestionId(topic.QuestionId, &subTopics); err != nil {
				log.Println(err)
				c.Data["json"] = Response{Status: "30022", Msg: "获取小题参数设置记录表失败  ", Data: err}
				return
			}
			subjectSubTopics = append(subjectSubTopics, subTopics...)

			testPapers := make([]model.TestPaper, 0)
			if err := model.FindTestPaperByQuestionId(topic.QuestionId, &testPapers); err != nil {
				log.Println(err)
				c.Data["json"] = Response{Status: "30022", Msg: "获取大题参数设置记录表失败  ", Data: err}
				return
			}
			subjectTestPapers = append(subjectTestPapers, testPapers...)
		}

		// Create a new sheet.
		index := f.NewSheet(subject.SubjectName)
		// Set value of a cell.
		f.SetCellValue(subject.SubjectName, "A1", "ticket_id")
		f.SetCellValue(subject.SubjectName, "B1", "name")
		f.SetCellValue(subject.SubjectName, "C1", "school")
		f.SetCellValue(subject.SubjectName, "D1", "mobile")
		// Set active sheet of the workbook.
		f.SetActiveSheet(index)

		var subjectTestPaperInfos []model.TestPaperInfo
		for i, subTopic := range subjectSubTopics {
			testPaperInfos := make([]model.TestPaperInfo, 0)
			model.FindTestPaperInfoByQuestionDetailId(subTopic.QuestionDetailId, &testPaperInfos)
			subjectTestPaperInfos = append(subjectTestPaperInfos, testPaperInfos...)

			f.SetCellValue(subject.SubjectName, string(byte(i+'E'))+"1", topicMap[subTopic.QuestionId]+"-"+subTopic.QuestionDetailName)
		}
		f.SetCellValue(subject.SubjectName, string(byte(len(subjectSubTopics)+'E'))+"1", "总分")

		for i := 2; i <= len(subjectTestPaperInfos)/len(subjectSubTopics)+1; i++ {
			f.SetCellValue(subject.SubjectName, "A"+strconv.Itoa(i), subjectTestPapers[(i-2)*len(topics)].TicketId)
			f.SetCellValue(subject.SubjectName, "B"+strconv.Itoa(i), subjectTestPapers[(i-2)*len(topics)].Candidate)
			f.SetCellValue(subject.SubjectName, "C"+strconv.Itoa(i), subjectTestPapers[(i-2)*len(topics)].School)
			f.SetCellValue(subject.SubjectName, "D"+strconv.Itoa(i), subjectTestPapers[(i-2)*len(topics)].Mobile)
			// 获取该用户的小题成绩
			infos, err := model.FindTestPaperInfoByTicketId(subjectTestPapers[(i-2)*len(topics)].TicketId)
			if err != nil {
				log.Println(err)
				c.Data["json"] = Response{Status: "30022", Msg: "获取小题参数设置记录表失败  ", Data: err}
				return
			}
			var sum int64
			for j, info := range infos {
				f.SetCellValue(subject.SubjectName, string(byte(j+'E'))+strconv.Itoa(i), info.FinalScore)
				sum += info.FinalScore
			}
			f.SetCellValue(subject.SubjectName, string(byte(len(subjectSubTopics)+'E'))+strconv.Itoa(i), sum)
		}

	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("../scores.xlsx"); err != nil {
		log.Println(err)
		c.Data["json"] = Response{Status: "30000", Msg: "excel 表导出错误", Data: err}
		return
	}

	c.Ctx.Output.Download("../scores.xlsx", "scores.xlsx")
}

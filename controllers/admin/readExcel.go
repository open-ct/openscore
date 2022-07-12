package admin

import (
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	. "openscore/controllers"
	"openscore/model"
	"os"
	"strconv"
	"strings"
)

/**
2.试卷导入
*/
func (c *AdminApiController) ReadExcel() {
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

				src := UploadPic(row[0]+rows[0][8+index], content)

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

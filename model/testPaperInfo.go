package model

import (
	"fmt"
	"log"

	"xorm.io/builder"
)

type TestPaperInfo struct {
	TestDetailId            int64  `json:"test_detail_id" xorm:"pk autoincr"`
	QuestionDetailId        int64  `json:"question_detail_id"`
	TestId                  int64  `json:"test_id"`
	PicSrc                  string `json:"pic_src"`
	ExaminerFirstId         int64  `json:"examiner_first_id" xorm:"default(-1)"`
	ExaminerFirstScore      int64  `json:"examiner_first_score" xorm:"default(-1)"`
	ExaminerFirstSelfScore  int64  `json:"examiner_first_self_score" xorm:"default(-1)"`
	ExaminerSecondId        int64  `json:"examiner_second_id" xorm:"default(-1)"`
	ExaminerSecondScore     int64  `json:"examiner_second_score" xorm:"default(-1)"`
	ExaminerSecondSelfScore int64  `json:"examiner_second_self_score" xorm:"default(-1)"`
	ExaminerThirdId         int64  `json:"examiner_third_id" xorm:"default(-1)"`
	ExaminerThirdScore      int64  `json:"examiner_third_score" xorm:"default(-1)"`
	ExaminerThirdSelfScore  int64  `json:"examiner_third_self_score" xorm:"default(-1)"`
	LeaderId                int64  `json:"leader_id" xorm:"default(-1)"`
	LeaderScore             int64  `json:"leader_score" xorm:"default(-1)"`
	FinalScore              int64  `json:"finale_score" xorm:"default(-1)"`
	FinalScoreId            int64  `json:"final_score_id" xorm:"default(-1)"`
	TicketId                string `json:"ticket_id"`
}

func (t *TestPaperInfo) GetTestPaperInfoByTestIdAndQuestionDetailId(testId int64, questionDetailId int64) error {
	has, err := adapter.Where("question_detail_id = ? and test_id = ?", questionDetailId, testId).Get(t)
	if !has || err != nil {
		log.Println("could not specific info")
	}
	return err
}

func (t *TestPaperInfo) GetTestPaperInfo(id int64) error {
	has, err := adapter.Where(builder.Eq{"test_detail_id": id}).Get(t)
	if !has && err != nil {
		log.Println("could not find test paper info")
		log.Println(err)
	}
	return err
}

func (t *TestPaperInfo) Update() error {
	code, err := adapter.Where(builder.Eq{"test_detail_id": t.TestDetailId}).AllCols().Update(t)
	if code == 0 && err != nil {
		log.Println("update test paper info fail")
		log.Println(err)
	}
	return err
}
func (t *TestPaperInfo) Insert() error {
	code, err := adapter.Insert(t)
	if code == 0 || err != nil {
		log.Println("Insert test paper info fail")
		log.Println(err)
	}
	return err
}

func GetTestInfoListByTestId(id int64, as *[]TestPaperInfo) error {
	err := adapter.Where("test_id = ?", id).Find(as)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

func GetTestInfoPicListByTestId(id int64, as *[]string) error {
	err := adapter.Table("test_paper_info").Select("pic_src").Where("test_id = ?", id).Find(as)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

func FindTestPaperInfoByQuestionDetailId(questionDetailId int64, t *[]TestPaperInfo) error {
	err := adapter.Where("question_detail_id = ?", questionDetailId).Find(t)
	if err != nil {
		log.Println("could not FindTestPaperInfoByQuestionId ")
		log.Println(err)
	}
	return err
}

func FindTestPaperInfoByTicketId(ticketId string) ([]*TestPaperInfo, error) {
	var infos []*TestPaperInfo
	err := adapter.Where("ticket_id = ?", ticketId).Find(&infos)
	fmt.Println("ticketId: ", ticketId, len(infos))

	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (t *TestPaperInfo) Delete() error {
	_, err := adapter.Where(builder.Eq{"test_id": t.TestId}).Delete(t)
	if err != nil {
		log.Println("delete fail")
	}
	return err
}

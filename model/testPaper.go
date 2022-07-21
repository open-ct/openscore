package model

import (
	"log"

	"xorm.io/builder"
)

type TestPaper struct {
	TestId                  int64  `json:"test_id" xorm:"pk autoincr"`
	QuestionId              int64  `json:"question_id"`
	Candidate               string `json:"candidate"`
	QuestionStatus          int64  `json:"question_status"`
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
	FinalScore              int64  `json:"final_score" xorm:"default(-1)" `
	FinalScoreId            int64  `json:"finale_score_id"`
	PracticeError           int64  `json:"practice_error"`
	CorrectingStatus        int64  `json:"correcting_status"`
	Mobile                  string `json:"mobile"`
	IsParent                int64  `json:"is_parent"`
	ClientIp                string `json:"client_ip"`
	Tag                     string `json:"tag"`
	School                  string `json:"school"`
	TicketId                string `json:"ticket_id"`
}

func (t *TestPaper) GetTestPaperByQuestionIdAndQuestionStatus(questionId int64, questionStatue int64) error {
	has, err := x.Where("question_id = ? and question_status = ?", questionId, questionStatue).Get(t)
	if !has || err != nil {
		log.Println("could not specific test")
	}
	return err
}

func (t *TestPaper) GetTestPaperByTestId(testId int64) error {
	has, err := x.Where("test_id = ?", testId).Get(t)
	if !has || err != nil {
		log.Println("could not GetTestPaperByTestId")
	}
	return err
}

func GetTestPaperListByQuestionIdAndQuestionStatus(questionId int64, questionStatue int64, tl *[]TestPaper) error {
	err := x.Where("question_id = ? and question_status = ?", questionId, questionStatue).Find(tl)
	if err != nil {
		log.Println("could not specific test")
		log.Println(err)
	}
	return err
}

func (t *TestPaper) GetTestPaper(id int64) (bool, error) {
	has, err := x.Where(builder.Eq{"test_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find test paper")
	}
	return has, err
}

func (t *TestPaper) Update() error {
	code, err := x.Where(builder.Eq{"test_id": t.TestId}).Update(t)
	if code == 0 || err != nil {
		log.Println("update test paper fail")
		log.Printf("%+v", err)
	}
	return err
}

func (t *TestPaper) Insert() (int64, error) {
	code, err := x.Insert(t)
	if code == 0 || err != nil {
		log.Println("insert test paper fail")
		log.Printf("%+v", err)
	}
	return t.TestId, err
}

func FindTestPaperByQuestionId(questionId int64, t *[]TestPaper) error {
	err := x.Where("question_id = ?", questionId).Find(t)
	if err != nil {
		log.Println("could not FindTestPaperByQuestionId ")
		log.Println(err)
	}
	return err
}
func FindTestPapersByTestId(testId, t *[]TestPaper) error {
	err := x.Where("question_id = ?", testId).Find(t)
	if err != nil {
		log.Println("could not FindTestPaperByQuestionId ")
		log.Println(err)
	}
	return err
}
func FindUnDistributeTest(id int64, t *[]TestPaper) error {
	err := x.Where("correcting_status = 0 AND question_id = ?", id).Find(t)
	if err != nil {
		log.Println("could not GetUnDistributeTest")
	}
	return err
}

func CountTestDistributionNumberByQuestionId(questionId int64) (count int64, err error) {
	testPaper := new(TestPaper)
	count, err1 := x.Where("question_id = ?", questionId).Where("correcting_status=?", 1).Count(testPaper)
	if err != nil {
		log.Println("CountTestDistributionNumberByQuestionId err ")
	}
	return count, err1
}
func CountFailTestNumberByUserId(userId int64, questionId int64) (count int64, err error) {
	testPaper := new(TestPaper)
	count, err1 := x.Where("question_id = ?", questionId).Where("examiner_first_id=? or examiner_second_id=?", userId, userId).Where("question_status=2 or question_status=3 ").Count(testPaper)
	if err != nil {
		log.Println("CountFailTestNumberByUserId err ")
	}
	return count, err1
}

func DeleteAllTest(questionId int64) error {
	_, err := x.Delete(&TestPaper{QuestionId: questionId})
	if err != nil {
		log.Println("delete fail")
	}
	return err
}

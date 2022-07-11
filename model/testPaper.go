package model

import (
	"log"

	"xorm.io/builder"
)

type TestPaper struct {
	Test_id                    int64  `json:"test_id" xorm:"pk autoincr"`
	Question_id                int64  `json:"question_id"`
	Candidate                  string `json:"candidate"`
	Question_status            int64  `json:"question_status"`
	Examiner_first_id          string `json:"examiner_first_id" xorm:"default('-1')"`
	Examiner_first_score       int64  `json:"examiner_first_score" xorm:"default(-1)"`
	Examiner_first_self_score  int64  `json:"examiner_first_self_score" xorm:"default(-1)"`
	Examiner_second_id         string `json:"examiner_second_id" xorm:"default('-1')"`
	Examiner_second_score      int64  `json:"examiner_second_score" xorm:"default(-1)"`
	Examiner_second_self_score int64  `json:"examiner_seconde_self_score" xorm:"default(-1)"`
	Examiner_third_id          string `json:"examiner_third_id" xorm:"default('-1')"`
	Examiner_third_score       int64  `json:"examiner_third_score" xorm:"default(-1)"`
	Examiner_third_self_score  int64  `json:"examiner_third_self_score" xorm:"default(-1)"`
	Leader_id                  string `json:"leader_id" xorm:"default('-1')"`
	Leader_score               int64  `json:"leader_score" xorm:"default(-1)"`
	Final_score                int64  `json:"final_score" xorm:"default(-1)" `
	Final_score_id             string `json:"finale_score_id"`
	Pratice_error              int64  `json:"pratice_error"`
	Correcting_status          int64  `json:"correcting_status"`
	Mobile                     string `json:"mobile"`
	IsParent                   int64  `json:"is_parent"`
	ClientIp                   string `json:"client_ip"`
	Tag                        string `json:"tag"`
	School                     string `json:"school"`
	TicketId                   string `json:"ticket_id"`
}

func (t *TestPaper) GetTestPaperByQuestionIdAndQuestionStatus(question_id int64, question_statue int64) error {
	has, err := x.Where("question_id = ? and question_status = ?", question_id, question_statue).Get(t)
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

func GetTestPaperListByQuestionIdAndQuestionStatus(question_id int64, question_statue int64, tl *[]TestPaper) error {
	err := x.Where("question_id = ? and question_status = ?", question_id, question_statue).Find(tl)
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
	code, err := x.Where(builder.Eq{"test_id": t.Test_id}).Update(t)
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
	return t.Test_id, err
}

func FindTestPaperByQuestionId(question_id int64, t *[]TestPaper) error {
	err := x.Where("question_id = ?", question_id).Find(t)
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
	err := x.Where("question_id=?", id).Where("correcting_status=?", 0).Find(t)
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
func CountFailTestNumberByUserId(userId string, questionId int64) (count int64, err error) {
	testPaper := new(TestPaper)
	count, err1 := x.Where("question_id = ?", questionId).Where("examiner_first_id=? or examiner_second_id=?", userId, userId).Where("question_status=2 or question_status=3 ").Count(testPaper)
	if err != nil {
		log.Println("CountFailTestNumberByUserId err ")
	}
	return count, err1
}

func DeleteAllTest(questionId int64) error {
	_, err := x.Delete(&TestPaper{Question_id: questionId})
	if err != nil {
		log.Println("delete fail")
	}
	return err
}

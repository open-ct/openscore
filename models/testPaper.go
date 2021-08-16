package models

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
	Examiner_first_score       int64  `json:"examiner_first_score"`
	Examiner_first_self_score  int64  `json:"examiner_first_self_score"`
	Examiner_second_id         string `json:"examiner_second_id" xorm:"default('-1')"`
	Examiner_second_score      int64  `json:"examiner_second_score"`
	Examiner_second_self_score int64  `json:"examiner_seconde_self_score"`
	Examiner_third_id          string `json:"examiner_third_id" xorm:"default('-1')"`
	Examiner_third_score       int64  `json:"examiner_third_score"`
	Examiner_third_self_score  int64  `json:"examiner_third_self_score"`
	Leader_id                  string `json:"leader_id" xorm:"default('-1')"`
	Leader_score               int64  `json:"leader_score"`
	Final_score                int64  `json:"final_score"`
	Final_score_id             string `json:"finale_score_id"`
	Pratice_error              int64  `json:"pratice_error"`
}

func (t *TestPaper) GetTestPaperByQuestionIdAndQuestionStatus(question_id int64, question_statue int64) error {
	has, err := x.Where("question_id = ? and question_status = ?", question_id, question_statue).Get(t)
	if !has || err != nil {
		log.Println("could not specific test")
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

func (t *TestPaper) GetTestPaper(id int64) error {
	has, err := x.Where(builder.Eq{"test_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find test paper")
	}
	return err
}

func (t *TestPaper) Update() error {
	code, err := x.Where(builder.Eq{"test_id": t.Test_id}).Update(t)
	if code == 0 || err != nil {
		log.Println("update test paper fail")
		log.Printf("%+v", err)
	}
	return err
}

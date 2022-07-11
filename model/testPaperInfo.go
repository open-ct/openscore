package model

import (
	"fmt"
	"log"

	"xorm.io/builder"
)

type TestPaperInfo struct {
	Test_detail_id             int64  `json:"test_detail_id" xorm:"pk autoincr"`
	Question_detail_id         int64  `json:"question_detail_id"`
	Test_id                    int64  `json:"test_id"`
	Pic_src                    string `json:"pic_src"`
	Examiner_first_id          string `json:"examiner_first_id" xorm:"default('-1')"`
	Examiner_first_score       int64  `json:"examiner_first_score" xorm:"default(-1)"`
	Examiner_first_self_score  int64  `json:"examiner_first_self_score" xorm:"default(-1)"`
	Examiner_second_id         string `json:"examiner_second_id" xorm:"default('-1')"`
	Examiner_second_score      int64  `json:"examiner_second_score" xorm:"default(-1)"`
	Examiner_second_self_score int64  `json:"examiner_second_self_score" xorm:"default(-1)"`
	Examiner_third_id          string `json:"examiner_third_id" xorm:"default('-1')"`
	Examiner_third_score       int64  `json:"examiner_third_score" xorm:"default(-1)"`
	Examiner_third_self_score  int64  `json:"examiner_third_self_score" xorm:"default(-1)"`
	Leader_id                  string `json:"leader_id" xorm:"default('-1')"`
	Leader_score               int64  `json:"leader_score" xorm:"default(-1)"`
	Final_score                int64  `json:"finale_score" xorm:"default(-1)"`
	Final_score_id             string `json:"final_score_id" xorm:"default('-1')"`
}

func (t *TestPaperInfo) GetTestPaperInfoByTestIdAndQuestionDetailId(testId int64, questionDetailId int64) error {
	fmt.Println("testId: ", testId, questionDetailId)

	has, err := x.Where("question_detail_id = ? and test_id = ?", questionDetailId, testId).Get(t)
	if !has || err != nil {
		log.Println("could not specific info")
	}
	return err
}

func (t *TestPaperInfo) GetTestPaperInfo(id int64) error {
	has, err := x.Where(builder.Eq{"test_detail_id": id}).Get(t)
	if !has && err != nil {
		log.Println("could not find test paper info")
		log.Println(err)
	}
	return err
}

func (t *TestPaperInfo) Update() error {
	code, err := x.Where(builder.Eq{"test_detail_id": t.Test_detail_id}).AllCols().Update(t)
	if code == 0 && err != nil {
		log.Println("update test paper info fail")
		log.Println(err)
	}
	return err
}
func (t *TestPaperInfo) Insert() error {
	code, err := x.Insert(t)
	if code == 0 || err != nil {
		log.Println("Insert test paper info fail")
		log.Println(err)
	}
	return err
}

func GetTestInfoListByTestId(id int64, as *[]TestPaperInfo) error {
	err := x.Where("test_id = ?", id).Find(as)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

func GetTestInfoPicListByTestId(id int64, as *[]string) error {
	err := x.Table("test_paper_info").Select("pic_src").Where("test_id = ?", id).Find(as)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

func FindTestPaperInfoByQuestionDetailId(questionDetailId int64, t *[]TestPaperInfo) error {
	err := x.Where("question_detail_id = ?", questionDetailId).Find(t)
	if err != nil {
		log.Println("could not FindTestPaperInfoByQuestionId ")
		log.Println(err)
	}
	return err
}
func (u *TestPaperInfo) Delete() error {
	_, err := x.Where(builder.Eq{"test_id": u.Test_id}).Delete(u)
	if err != nil {
		log.Println("delete fail")
	}
	return err
}

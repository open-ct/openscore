package models

import (
	"log"

	"xorm.io/builder"
)

type UnderCorrectedPaper struct {
	UnderCorrected_id  int64  `json:"underCorrected_id" xorm:"pk autoincr"`
	User_id            string `json:"user_id"`
	Test_id            int64  `json:"test_id"`
	Question_id        int64  `json:"question_id"`
	Test_question_type int64  `json:"test_question_type"`
	Problem_type       int64  `json:"problem_type" xorm:"default(-1)"`
}

func (u *UnderCorrectedPaper) GetUnderCorrectedPaper(userId string, testId int64) error {
	has, err := x.Where(builder.Eq{"test_id": testId, "user_id": userId}).Get(u)
	if !has || err != nil {
		log.Println("could not find under corrected paper")
		log.Println(err)
	}
	return err
}

func (u *UnderCorrectedPaper) Delete() error {
	code, err := x.Where(builder.Eq{"test_id": u.Test_id, "user_id": u.User_id}).Delete(u)
	if code == 0 || err != nil {
		log.Println("delete fail")
	}
	return err
}

func (u *UnderCorrectedPaper) Save() error {
	code, err := x.Insert(u)
	if code == 0 || err != nil {
		log.Println("insert paper fail")
		log.Println(err)
	}
	return err
}

func (u *UnderCorrectedPaper) IsDuplicate() (bool, error) {
	var temp UnderCorrectedPaper
	has, err := x.Where(builder.Eq{"test_id": u.Test_id, "problem_type": u.Problem_type}).Get(&temp)
	if !has || err != nil {
		log.Println(err)
	}
	return has, err
}

func GetDistributedPaperByUserId(id string, up *[]UnderCorrectedPaper) error {
	err := x.Where("user_id = ?", id).Find(up)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

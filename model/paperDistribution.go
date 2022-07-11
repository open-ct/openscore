package model

import (
	"log"

	"xorm.io/builder"
)

type PaperDistribution struct {
	Distribution_id          int64  `json:"distribution_id" xorm:"pk autoincr"`
	User_id                  string `json:"user_id"`
	Question_id              int64  `json:"question_id"`
	Test_distribution_number int64  `json:"test_distribution_number"`
	PaperType                int64  `json:"paperType"`
}

func (u *PaperDistribution) GetPaperDistribution(id string) error {
	has, err := x.Where(builder.Eq{"user_id": id}).Get(u)
	if !has || err != nil {
		log.Println("could not find paper distribution")
	}
	return err
}
func FindPaperDistributionByQuestionId(paperDistributions *[]PaperDistribution, questionId int64) error {
	err := x.Where("question_id = ?", questionId).Find(paperDistributions)
	if err != nil {
		log.Println("FindPaperDistributionByQuestionId err ")
	}
	return err
}

func (u *PaperDistribution) Save() error {
	code, err := x.Insert(u)
	if code == 0 || err != nil {
		log.Println("insert PaperDistribution fail")
		log.Println(err)
	}
	return err
}
func CountUserDistributionNumberByQuestionId(questionId int64) (count int64, err error) {
	paperDistribution := new(PaperDistribution)
	count, err1 := x.Where("question_id = ?", questionId).Count(paperDistribution)
	if err != nil {
		log.Println("countUserDistributionNumberByQuestionId err ")
	}
	return count, err1
}

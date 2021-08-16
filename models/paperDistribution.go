package models

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

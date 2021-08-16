package models

import (
	"log"

	"xorm.io/builder"
)

type SubTopic struct {
	Question_detail_id    int64  `json:"question_detail_id" xorm:"pk autoincr" `
	Question_detail_name  string `json:"question_detail_name"`
	Question_id           int64  `json:"question_id"`
	Question_detail_score int64  `json:"question_detail_score"`
	Score_type            string `json:"score_type"`
}

func GetSubTopicsByQuestionId(id int64, st *[]SubTopic) error {
	err := x.Where("question_id = ?", id).Find(st)
	if err != nil {
		log.Println("could not find any SubTopic")
	}
	return err
}

func GetSubTopicsByTestId(id int64, st *[]SubTopic) error {
	err := x.Where(builder.Eq{"question_id": id}).Find(st)
	if err != nil {
		log.Println("could not find any SubTopic")
		log.Println(err)
	}
	return err
}

func (st *SubTopic) GetSubTopic(id int64) error {
	has, err := x.Where(builder.Eq{"question_detail_id": id}).Get(st)
	if !has || err != nil {
		log.Println("could not find SubTopic")
	}
	return err
}

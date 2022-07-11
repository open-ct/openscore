package model

import (
	"log"

	"xorm.io/builder"
)

type SubTopic struct {
	QuestionDetailId    int64  `json:"question_detail_id" xorm:"pk autoincr"`
	QuestionDetailName  string `json:"question_detail_name"`
	QuestionId          int64  `json:"question_id"`
	QuestionDetailScore int64  `json:"question_detail_score"`
	ScoreType           string `json:"score_type"`
}

func FindSubTopicsByQuestionId(id int64, st *[]SubTopic) error {
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
func InsertSubTopic(subTopic *SubTopic) (err1 error, questionDetailId int64) {
	_, err := x.Insert(subTopic)
	if err != nil {
		log.Println("GetTopicList err ")
	}

	return err, subTopic.QuestionDetailId
}

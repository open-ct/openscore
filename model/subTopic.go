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
	IsSecondScore       bool   `json:"is_second_score"`
}

func FindSubTopicsByQuestionId(id int64, st *[]SubTopic) error {
	err := adapter.engine.Where("question_id = ?", id).Find(st)
	if err != nil {
		log.Println("could not find any SubTopic")
	}
	return err
}

func GetSubTopicsByTestId(id int64, st *[]SubTopic) error {
	err := adapter.engine.Where(builder.Eq{"question_id": id}).Find(st)
	if err != nil {
		log.Println("could not find any SubTopic")
		log.Println(err)
	}
	return err
}

func (s *SubTopic) GetSubTopic(id int64) error {
	has, err := adapter.engine.Where(builder.Eq{"question_detail_id": id}).Get(s)
	if !has || err != nil {
		log.Println("could not find SubTopic")
	}
	return err
}

func InsertSubTopic(subTopic *SubTopic) error {
	_, err := adapter.engine.Insert(subTopic)
	if err != nil {
		log.Println("GetTopicList err ")
	}

	return err
}

func (s *SubTopic) Delete() error {
	_, err := adapter.engine.Where("question_detail_id = ?", s.QuestionDetailId).Delete(s)
	return err
}

func (s *SubTopic) Update() error {
	code, err := adapter.engine.Where(builder.Eq{"question_detail_id": s.QuestionDetailId}).Update(s)
	if code == 0 || err != nil {
		log.Println("update subTopic fail")
		log.Printf("%+v", err)
	}
	return err
}

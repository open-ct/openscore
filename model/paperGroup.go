package model

import (
	"errors"
	"log"
	"xorm.io/builder"
)

type PaperGroup struct {
	Id         int64   `json:"id" xorm:"pk autoincr"`
	GroupName  string  `json:"group_name"`
	TestIds    []int64 `json:"test_ids"`
	QuestionId int64   `json:"question_id"`
}

func (p *PaperGroup) Update() error {
	code, err := adapter.engine.Where(builder.Eq{"id": p.Id}).Update(p)
	if code == 0 || err != nil {
		log.Println("update PaperGroup fail")
		log.Printf("%+v", err)
	}
	return err
}

func CreatePaperGroup(group *PaperGroup) error {
	_, err := adapter.engine.Insert(group)
	if err != nil {
		log.Println("GetTopicList err ")
	}

	return err
}

func GetGroupByGroupId(id int64) (*PaperGroup, error) {
	var group PaperGroup
	get, err := adapter.engine.Where("id=?", id).Get(&group)

	if err != nil || !get {
		log.Println("GetGroupByGroupId err ")
		return nil, errors.New("GetGroupByGroupId")
	}

	return &group, nil
}

func GetGroupThanLastId(id int64, questionId int64) (*PaperGroup, bool, error) {
	var group PaperGroup
	get, err := adapter.engine.Where("id > ?", id).Where("question_id = ?", questionId).Get(&group)
	if err != nil {
		log.Println("GetGroupByGroupId err ")
		return nil, false, errors.New("GetGroupByGroupId")
	}

	return &group, get, nil
}

func ListPaperGroup() ([]*PaperGroup, error) {
	var groups []*PaperGroup
	err := adapter.engine.Find(&groups)
	if err != nil {
		log.Println("ListPaperGroup err ")
		return nil, err
	}

	return groups, nil
}

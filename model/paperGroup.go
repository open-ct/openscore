package model

import "log"

type PaperGroup struct {
	Id        int64   `json:"id" xorm:"pk autoincr"`
	GroupName string  `json:"group_name"`
	TestIds   []int64 `json:"test_ids"`
}

func CreatePaperGroup(group *PaperGroup) error {
	_, err := adapter.engine.Insert(group)
	if err != nil {
		log.Println("GetTopicList err ")
	}

	return err
}

func ListPaperGroup() ([]*PaperGroup, error) {
	var groups []*PaperGroup
	err := adapter.engine.Find(&groups)
	if err != nil {
		log.Println("FindSubjectList err ")
		return nil, err
	}
	return groups, nil
}

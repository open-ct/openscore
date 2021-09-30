package models

import (
	"log"
	"time"

	"xorm.io/builder"
)

// Author: Junlang
// struct : Topic(大题)
// comment: must capitalize the first letter of the field in Topic
type Topic struct {
	Question_id    int64     `json:"question_id" xorm:"pk autoincr"`
	Question_name  string    `json:"question_name" xorm:"varchar(50)"`
	Subject_name   string    `json:"subject_name" xorm:"varchar(50)"`
	Standard_error int64     `json:"standard_error"`
	Question_score int64     `json:"quetsion_score"`
	Score_type     int64     `json:"score_type"`
	Import_number  int64     `json:"import_number"`
	Import_time    time.Time `json:"import_time"`
	Subject_Id    int64  	 `json:"subject_id"`
	SelfScoreRate  int64     `json:"self_score_rate"`
}

func (t *Topic) GetTopic(id int64) error {
	has, err := x.Where(builder.Eq{"question_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find topic")
	}
	return err
}


func GetTopicList ( topics *[]Topic) error{
	err := x.Find(topics)
	if err!=nil {
		log.Println("GetTopicList err ")
	}
	return  err
}
func FindTopicBySubNameList ( topics *[]Topic,subjectName string) error{
	err := x.Where("subject_name=?",subjectName).Find(topics)
	if err!=nil {
		log.Println("FindTopicBySubNameList err ")
	}
	return  err
}
func FindTopicList ( topics *[]Topic) error{
	err := x.Find(topics)
	if err!=nil {
		log.Println("FindTopicList err ")
	}
	return  err
}
func InsertTopic ( topic *Topic)(err1 error,questionId int64) {
	_,err:= x.Insert(topic)
	if err!=nil {
		log.Println("GetTopicList err ")
	}

	return  err,topic.Question_id
}
//func Update ( topic *Topic,questionId int64)error {
//	_,err:= x.Where("question_id=?",questionId).Update(&topic)
//	if err!=nil {
//		log.Println("Update topic err ")
//	}
//
//	return  err
//}
func (t *Topic) Update() error {
	code, err := x.Where(builder.Eq{"question_id": t.Question_id}).Update(t)
	if code == 0 || err != nil {
		log.Println("update Topic paper fail")
		log.Printf("%+v", err)
	}
	return err
}
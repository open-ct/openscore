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
}

func (t *Topic) GetTopic(id int64) error {
	has, err := x.Where(builder.Eq{"question_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find topic")
	}
	return err
}

func GetDistributedTestIdPaperByUserId(id string, up *[]int64) error {
	err := x.Table("under_corrected_paper").Select("test_id").Where("user_id = ?", id).Find(up)
	if err != nil {
		log.Panic(err)
		log.Println("could not find any paper")
	}
	return err
}

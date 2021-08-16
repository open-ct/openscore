package models

import (
	"log"
	"time"

	"xorm.io/builder"
)

type ScoreRecord struct {
	Record_id        int64     `json:"record_id" xorm:"pk autoincr"`
	Question_id      int64     `json:"question_id"`
	Test_id          int64     `json:"test_id"`
	User_id          string    `json:"user_id"`
	Score_time       time.Time `json:"score_time"`
	Score            int64     `json:"score"`
	Test_record_type int64     `json:"test_record_type"`
	Problem_type     int64     `json:"problem_type" xorm:"default(-1)"`
}

func (s *ScoreRecord) GetTopic(id int64) error {
	has, err := x.Where(builder.Eq{"Question_id": id}).Get(s)
	if !has || err != nil {
		log.Println("could not find user")
	}
	return err
}

func (r *ScoreRecord) Save() error {
	code, err := x.Insert(r)
	if code == 0 || err != nil {
		log.Println("insert record fail")
	}
	return err
}

func GetLatestRecords(userId string, records *[]ScoreRecord) error {
	err := x.Limit(10).Table("score_record").Select("test_id, score, score_time").Where(builder.Eq{"user_id": userId}).Desc("record_id").Find(records)
	if err != nil {
		log.Panic(err)
		log.Println("could not find any paper")
	}

	return err
}

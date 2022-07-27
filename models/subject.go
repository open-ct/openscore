package models

import (
	"log"
)

type Subject struct {
	SubjectId   int64  `json:"subjectId" xorm:"pk autoincr"`
	SubjectName string `json:"subject_name"`
}

func FindSubjectList(subjects *[]Subject) error {
	err := adapter.engine.Find(subjects)
	if err != nil {
		log.Println("FindSubjectList err ")
	}
	return err
}

func InsertSubject(subject *Subject) (err1 error, questionId int64) {
	_, err := adapter.engine.Insert(subject)
	if err != nil {
		log.Println("GetTopicList err ")
	}

	return err, subject.SubjectId
}
func GetSubjectBySubjectName(subject *Subject, subjectName string) (bool, error) {
	get, err := adapter.engine.Where("subject_name=?", subjectName).Get(subject)
	if err != nil {
		log.Println("FindSubjectList err ")
	}
	return get, err
}

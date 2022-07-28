package model

import (
	"log"
)

type Subject struct {
	SubjectId   int64  `json:"subject_id" xorm:"pk autoincr"`
	SubjectName string `json:"subject_name"`
}

func FindSubjectList(subjects *[]Subject) error {
	err := adapter.Find(subjects)
	if err != nil {
		log.Println("FindSubjectList err ")
	}
	return err
}

func InsertSubject(subject *Subject) (err1 error, questionId int64) {
	_, err := adapter.Insert(subject)
	if err != nil {
		log.Println("GetTopicList err ")
	}

	return err, subject.SubjectId
}

func GetSubjectBySubjectName(subject *Subject, subjectName string) (bool, error) {
	get, err := adapter.Where("subject_name=?", subjectName).Get(subject)
	if err != nil {
		log.Println("FindSubjectList err ")
	}
	return get, err
}

func GetSubjectById(id int64) (string, error) {
	subject := &Subject{}
	ok, err := adapter.Where("subject_id = ?", id).Get(subject)
	if err != nil || !ok {
		log.Println("GetSubjectById err")
	}

	return subject.SubjectName, err
}

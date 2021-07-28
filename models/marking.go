package models

import (
	"log"
	"time"
)

// Author: Junlang
// struct : Topic(大题)
// comment: must capitalize the first letter of the field in Topic
type Topic struct {
	Question_id    int    `xorm:"id pk"`
	Question_name  string `xorm:"varchar(50)"`
	Subject_name   string `xorm:"varchar(50)"`
	Standard_error int
	Question_score int
	Score_type     int
	Import_number  int
	Import_time    time.Time `xorm:updated`
}

type SubTopic struct {
	Question_detail_id    int `xorm:"id pk" `
	Question_detail_name  string
	Question_id           int
	Question_detail_score int
}

type TestPaper struct {
	Test_id                    int
	Question_id                int
	Candidate                  string
	Correcting_status          int
	Question_status            int
	Examiner_first_id          int
	Examiner_first_score       int
	Examiner_first_self_score  int
	Examiner_second_id         int
	Examiner_second_score      int
	Examiner_second_self_score int
	Examiner_third_id          int
	Examiner_third_score       int
	Examiner_third_self_score  int
	Leader_id                  int
	Leader_score               int
	Final_score                int
	Problem_type               int
	Pratice_error              int
}

type TestPaperInfo struct {
	Test_detail_id             int
	Question_detail_id         int
	Test_id                    int
	Pic_src                    string
	Examiner_first_id          int
	Examiner_first_score       int
	Examiner_first_self_score  int
	Examiner_second_id         int
	Examiner_second_score      int
	Examiner_second_self_score int
	Examiner_third_id          int
	Examiner_third_score       int
	Examiner_third_self_score  int
	Leader_id                  int
	Leader_score               int
	Final_score                int
}

func initMarkingModels() {
	err := x.Sync2(new(Topic), new(SubTopic), new(TestPaper), new(TestPaperInfo))
	if err != nil {
		log.Println(err)
	}
}

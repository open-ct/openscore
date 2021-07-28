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
	Question_id    int64  `xorm:"id pk"`
	Question_name  string `xorm:"varchar(50)"`
	Subject_name   string `xorm:"varchar(50)"`
	Standard_error int64
	Question_score int64
	Score_type     int64
	Import_number  int64
	Import_time    time.Time `xorm:updated`
}

type SubTopic struct {
	Question_detail_id    int64 `xorm:"id pk" `
	Question_detail_name  string
	Question_id           int64
	Question_detail_score int64
}

type TestPaper struct {
	Test_id                    int64
	Question_id                int64
	Candidate                  string
	Correcting_status          int64
	Question_status            int64
	Examiner_first_id          int64
	Examiner_first_score       int64
	Examiner_first_self_score  int64
	Examiner_second_id         int64
	Examiner_second_score      int64
	Examiner_second_self_score int64
	Examiner_third_id          int64
	Examiner_third_score       int64
	Examiner_third_self_score  int64
	Leader_id                  int64
	Leader_score               int64
	Final_score                int64
	Problem_type               int64
	Pratice_error              int64
}

type TestPaperInfo struct {
	Test_detail_id             int64
	Question_detail_id         int64
	Test_id                    int64
	Pic_src                    string
	Examiner_first_id          int64
	Examiner_first_score       int64
	Examiner_first_self_score  int64
	Examiner_second_id         int64
	Examiner_second_score      int64
	Examiner_second_self_score int64
	Examiner_third_id          int64
	Examiner_third_score       int64
	Examiner_third_self_score  int64
	Leader_id                  int64
	Leader_score               int64
	Final_score                int64
}

type UnderCorrectedPaper struct {
	User_id            int64
	Test_id            int64
	Question_id        int64
	Test_question_type int64
}

type ScoreRecord struct {
	Record_id        int64
	Test_id          int64
	Tser_id          int64
	Score_time       int64
	Score            int64
	Self_score       int64
	Test_record_type int64
	Score_type       int64
}

type PaperDistribution struct {
	User_id                  int64
	Question_id              int64
	Test_distribution_number int64
	Test_success_number      int64
	Test_remaining_number    int64
	PaperType                int64
}

func initMarkingModels() {
	err := x.Sync2(new(Topic), new(SubTopic), new(TestPaper), new(TestPaperInfo))
	if err != nil {
		log.Println(err)
	}
}
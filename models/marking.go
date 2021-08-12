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
	Question_id    int64  `xorm:"pk autoincr"`
	Question_name  string `xorm:"varchar(50)"`
	Subject_name   string `xorm:"varchar(50)"`
	Standard_error int64
	Question_score int64
	Score_type     int64
	Import_number  int64
	Import_time    time.Time `xorm:updated`
}

type SubTopic struct {
	Question_detail_id    int64 `xorm:"pk autoincr" `
	Question_detail_name  string
	Question_id           int64
	Question_detail_score int64
}

type TestPaper struct {
	Test_id                    int64 `xorm:"pk autoincr"`
	Question_id                int64
	Candidate                  string
	Question_status            int64
	Examiner_first_id          string `xorm:"default('-1')"`
	Examiner_first_score       int64
	Examiner_first_self_score  int64
	Examiner_second_id         string `xorm:"default('-1')"`
	Examiner_second_score      int64
	Examiner_second_self_score int64
	Examiner_third_id          string `xorm:"default('-1')"`
	Examiner_third_score       int64
	Examiner_third_self_score  int64
	Leader_id                  string `xorm:"default('-1')"`
	Leader_score               int64
	Final_score                int64
	Final_score_id             string
	Pratice_error              int64
	Answer_test_id             int64
	Example_test_id            int64
}

type TestPaperInfo struct {
	Test_detail_id             int64 `xorm:"pk autoincr"`
	Question_detail_id         int64
	Test_id                    int64
	Pic_src                    string
	Examiner_first_id          string `xorm:"default('-1')"`
	Examiner_first_score       int64
	Examiner_first_self_score  int64
	Examiner_second_id         string `xorm:"default('-1')"`
	Examiner_second_score      int64
	Examiner_second_self_score int64
	Examiner_third_id          string `xorm:"default('-1')"`
	Examiner_third_score       int64
	Examiner_third_self_score  int64
	Leader_id                  string `xorm:"default('-1')"`
	Leader_score               int64
	Final_score                int64
	Final_score_id             string `xorm:"default('-1')"`
}

type UnderCorrectedPaper struct {
	UnderCorrected_id  int64 `xorm:"pk autoincr"`
	User_id            string
	Test_id            int64
	Question_id        int64
	Test_question_type int64
	Problem_type       int64 `xorm:"default(-1)"`
}

type ScoreRecord struct {
	Record_id        int64 `xorm:"pk autoincr"`
	Question_id      int64
	Test_id          int64
	User_id          string
	Score_time       time.Time
	Score            int64
	Test_record_type int64
	Problem_type     int64 `xorm:"default(-1)"`
}

type PaperDistribution struct {
	Distribution_id          int64 `xorm:"pk autoincr"`
	User_id                  string
	Question_id              int64
	Test_distribution_number int64
	PaperType                int64
}

func initMarkingModels() {
	err := x.Sync2(new(Topic), new(SubTopic), new(TestPaper), new(TestPaperInfo), new(ScoreRecord), new(UnderCorrectedPaper), new(PaperDistribution))
	if err != nil {
		log.Println(err)
	}
}

func (t *Topic) GetTopic(id int64) error {
	has, err := x.Where(builder.Eq{"question_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find topic")
	}
	return err
}

func GetSubTopicsByQuestionId(id int64, st *[]SubTopic) error {
	err := x.Where("question_id = ?", id).Find(st)
	if err != nil {
		log.Println("could not find any SubTopic")
	}
	return err
}

func GetSubTopicsByTestId(id int64, st *[]SubTopic) error {
	err := x.Where(builder.Eq{"question_id": id}).Find(st)
	if err != nil {
		log.Println("could not find any SubTopic")
		log.Println(err)
	}
	return err
}

func GetDistributedPaperByUserId(id int64, up *[]UnderCorrectedPaper) error {
	err := x.Where("user_id = ?", id).Find(up)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

func GetTestInfoListByTestId(id int64, as *[]TestPaperInfo) error {
	err := x.Where("test_id = ?", id).Find(as)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}

func (t *TestPaper) GetTestPaperByQuestionIdAndQuestionStatus(question_id int64, question_statue int64) error {
	has, err := x.Where("question_id = ? and question_status = ?", question_id, question_statue).Get(t)
	if !has || err != nil {
		log.Println("could not specific test")
	}
	return err
}

func GetTestPaperListByQuestionIdAndQuestionStatus(question_id int64, question_statue int64, tl *[]TestPaper) error {
	err := x.Where("question_id = ? and question_status = ?", question_id, question_statue).Find(tl)
	if err != nil {
		log.Println("could not specific test")
		log.Println(err)
	}
	return err
}

func (t *TestPaperInfo) GetTestPaperInfoByTestIdAndQuestionDetailId(testId int64, questionDetailId int64) error {
	has, err := x.Where("question_detail_id = ? and test_id = ?", questionDetailId, testId).Get(t)
	if !has || err != nil {
		log.Println("could not specific info")
	}
	return err
}

func (st *SubTopic) GetSubTopic(id int64) error {
	has, err := x.Where(builder.Eq{"question_detail_id": id}).Get(st)
	if !has || err != nil {
		log.Println("could not find SubTopic")
	}
	return err
}

func (t *TestPaper) GetTestPaper(id int64) error {
	has, err := x.Where(builder.Eq{"test_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find test paper")
	}
	return err
}

func (t *TestPaperInfo) GetTestPaperInfo(id int64) error {
	has, err := x.Where(builder.Eq{"test_detail_id": id}).Get(t)
	if !has || err != nil {
		log.Println("could not find test paper info")
		log.Println(err)
	}
	return err
}

func (u *UnderCorrectedPaper) GetUnderCorrectedPaper(userId string, testId int64) error {
	has, err := x.Where(builder.Eq{"test_id": testId, "user_id": userId}).Get(u)
	if !has || err != nil {
		log.Println("could not find under corrected paper")
		log.Println(err)
	}
	return err
}

func (u *UnderCorrectedPaper) Delete() error {
	code, err := x.Where(builder.Eq{"test_id": u.Test_id, "user_id": u.User_id}).Delete(u)
	if code == 0 || err != nil {
		log.Println("delete fail")
	}
	return err
}

func (u *PaperDistribution) GetPaperDistribution(id string) error {
	has, err := x.Where(builder.Eq{"user_id": id}).Get(u)
	if !has || err != nil {
		log.Println("could not find paper distribution")
	}
	return err
}

func (s *ScoreRecord) GetTopic(id int64) error {
	has, err := x.Where(builder.Eq{"Question_id": id}).Get(s)
	if !has || err != nil {
		log.Println("could not find user")
	}
	return err
}

func (t *TestPaperInfo) Update() error {
	code, err := x.Where(builder.Eq{"test_detail_id": t.Test_detail_id}).Update(t)
	if code == 0 || err != nil {
		log.Println("update test paper info fail")
		log.Println(err)
	}
	return err
}

func (t *TestPaper) Update() error {
	code, err := x.Where(builder.Eq{"test_id": t.Test_id}).Update(t)
	if code == 0 || err != nil {
		log.Println("update test paper fail")
		log.Printf("%+v", err)
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

func (u *UnderCorrectedPaper) Save() error {
	code, err := x.Insert(u)
	if code == 0 || err != nil {
		log.Println("insert paper fail")
		log.Println(err)
	}
	return err
}

func (u *UnderCorrectedPaper) IsDuplicate() (bool, error) {
	var temp UnderCorrectedPaper
	has, err := x.Where(builder.Eq{"test_id": u.Test_id, "problem_type": u.Problem_type}).Get(&temp)
	if !has || err != nil {
		log.Println(err)
	}
	return has, err
}

func GetLatestRecores(userId string, records *[]ScoreRecord) error {
	// x.QueryString("select top 10 * from scoreRecord where user_id = " + strconv.FormatInt(userId, 10) + " order by record_id desc")
	err := x.Limit(10).Where(builder.Eq{"user_id": userId}).Desc("record_id").Find(records)
	if err != nil {
		log.Println("could not find any paper")
	}

	return err
}

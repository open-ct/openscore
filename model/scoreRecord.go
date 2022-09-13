package model

import (
	"log"
	"time"

	"xorm.io/builder"
)

type ScoreRecord struct {
	RecordId       int64     `json:"record_id" xorm:"pk autoincr"`
	QuestionId     int64     `json:"question_id"`
	TestId         int64     `json:"test_id"`
	UserId         int64     `json:"user_id"`
	ScoreTime      time.Time `json:"score_time"`
	Score          int64     `json:"score"`
	TestRecordType int64     `json:"test_record_type"`
	ProblemType    int64     `json:"problem_type" xorm:"default(-1)"`
	TestFinish     int64     `json:"test_finish"`
}

func (s *ScoreRecord) GetTopic(id int64) error {
	has, err := adapter.engine.Where(builder.Eq{"Question_id": id}).Get(s)
	if !has || err != nil {
		log.Println("could not find user")
	}
	return err
}

func (s *ScoreRecord) GetRecordByTestId(testId int64, userId int64) error {
	has, err := adapter.engine.Where(builder.Eq{"test_id": testId}).Where("user_id=?", userId).Where("test_record_type !=0").Get(s)
	if !has || err != nil {
		log.Println("could not find ScoreRecord")
	}
	return err
}

func GetRecordByTestId(testId int64) (*ScoreRecord, error) {
	s := &ScoreRecord{}
	has, err := adapter.engine.Where(builder.Eq{"test_id": testId}).Where("test_record_type !=0").Get(s)
	if !has || err != nil {
		log.Println("could not find ScoreRecord")
		return nil, err
	}
	return s, nil
}

func (s *ScoreRecord) Save() error {
	code, err := adapter.engine.Insert(s)
	if code == 0 || err != nil {
		log.Println("insert record fail")
	}
	return err
}

func ListUserScoreRecord(userId int64) ([]ScoreRecord, error) {
	var records []ScoreRecord
	err := adapter.engine.Where("test_record_type=1").Where("user_id=?", userId).Find(&records)
	if err != nil {
		log.Println("ListUserScoreRecord err ")
	}
	return records, err
}

func GetLatestRecords(userId int64, records *[]ScoreRecord) error {
	err := adapter.engine.Limit(10).Table("score_record").Select("test_id, score, score_time").Where(builder.Eq{"user_id": userId}).Where("test_record_type=1 or test_record_type=2 or test_record_type=3").Desc("record_id").Find(records)
	if err != nil {
		log.Panic(err)
		log.Println("could not find any paper")
	}

	return err
}

func CountProblemFinishNumberByQuestionId(questionId int64) (count int64, err error) {
	record := new(ScoreRecord)
	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=?", 6).Count(record)
	if err != nil {
		log.Println("CountProblemFinishNumberByQuestionId err ")
	}
	return count, err1
}

func CountSelfScore(userId int64, questionId int64) (count int64, err error) {
	record := new(ScoreRecord)
	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=0").Where("user_id=?", userId).Count(record)
	if err != nil {
		log.Println("CountFailTestNumberByUserId err ")
	}
	return count, err1
}

// func CountFailTestNumberByUserId(userId int64,questionId int64)(count int64,err error) {
//	record :=new (ScoreRecord)
//	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=5").Where("user_id=?", userId).Count(record)
//	if err!=nil {
//		log.Println("CountFailTestNumberByUserId err ")
//	}
//	return count,err1
// }

func CountProblemNumberByQuestionId(questionId int64) (count int64) {
	record := new(ScoreRecord)
	count, err := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=?", 5).Count(record)
	if err != nil {
		log.Println("CountProblemNumberByQuestionId err ")
	}
	return count
}

func CountArbitramentFinishNumberByQuestionId(questionId int64) (count int64, err error) {
	record := new(ScoreRecord)
	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=?", 4).Count(record)
	if err != nil {
		log.Println("CountArbitramentFinishNumberByQuestionId err ")
	}
	return count, err1
}

// func CountFinishScoreNumberByUserId(userId int64,questionId int64)(count int64 ,err error) {
//	record :=new (ScoreRecord)
//	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("user_id =?",userId).Where("test_finish=1").Count(record)
//	if err1!=nil {
//		log.Println("CountFinishScoreNumberByUserId err ")
//	}
//	return count,err1
// }

func CountFinishScoreNumberByQuestionId(questionId int64) (count int64, err error) {
	record := new(ScoreRecord)
	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where(" test_record_type!=7 ").Where(" test_record_type!=0 ").Where("test_finish=1").Count(record)
	if err != nil {
		log.Println("CountFinishScoreNumberByQuestionId err ")
	}
	return count, err1
}

func FindFinishScoreByQuestionId(finishScores *[]ScoreRecord, questionId int64) error {
	err := adapter.engine.Where("question_id = ?", questionId).Where("test_finish=1").Find(*finishScores)
	if err != nil {
		log.Println("CountFinishScoreNumberByQuestionId err ")
	}
	return err
}

func FindFinishTestByUserId(scoreRecord *[]ScoreRecord, userId int64, questionId int64) error {
	err := adapter.engine.Where("question_id = ?", questionId).Where("user_id=?", userId).Where("test_record_type =1 or test_record_type =2 ").Where("test_finish=1").Find(scoreRecord)
	if err != nil {
		log.Println("FindFinishTestNumberByUserId err ")
	}
	return err
}

func CountFirstScoreNumberByQuestionId(questionId int64) (int64, error) {
	record := new(ScoreRecord)
	count, err := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=1 or test_record_type=2 ").Where("test_finish=1").Count(record)
	if err != nil {
		log.Println("CountFirstScoreNumberByQuestionId err ")
	}
	return count, err
}

func CountSecondScoreNumberByQuestionId(questionId int64) (count int64, err error) {
	record := new(ScoreRecord)
	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=1 or test_record_type=2 ").Where("test_finish=1").Count(record)
	if err != nil {
		log.Println("CountSecondScoreNumberByQuestionId err ")
	}
	return count, err1
}

func CountThirdScoreNumberByQuestionId(questionId int64) (count int64, err error) {
	record := new(ScoreRecord)
	count, err1 := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=1 or test_record_type=2 ").Where("test_finish=1").Count(record)
	if err != nil {
		log.Println("CountThirdScoreNumberByQuestionId err ")
	}
	return count, err1
}

func CountTestScoreNumberByUserId(userId int64, questionId int64) (count int64, err1 error) {
	record := new(ScoreRecord)
	count, err := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=1 or test_record_type=2 ").Where("user_id=?", userId).Count(record)
	if err != nil {
		log.Println("CountFinishTestNumberByUserId err ")
	}
	return count, err
}

func SumFinishScore(userId int64, questionId int64) (sum float64, err1 error) {
	record := new(ScoreRecord)
	sum, err := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=1 or test_record_type=2  ").Where("user_id=?", userId).Sum(record, "score")
	if err != nil {
		log.Println("SumFinishScore err ")
	}
	return sum, err
}

func AverageFinishScore(userId int64) (float64, error) {
	record := new(ScoreRecord)
	sum, err := adapter.engine.Where("test_record_type=1 or test_record_type=2  ").Where("user_id = ?", userId).Sum(record, "score")
	if err != nil {
		log.Println("SumFinishScore err ")
	}
	num, err := adapter.engine.Where("test_record_type=1 or test_record_type=2  ").Where("user_id = ?", userId).Count(record)

	if num == 0 {
		return 0, err
	}
	return sum / float64(num), err
}

func FindFinishScoreRecordListByQuestionId(scoreRecordList *[]ScoreRecord, questionId int64) error {
	err := adapter.engine.Where("question_id = ?", questionId).Where("test_record_type=1 or test_record_type=2  ").Find(scoreRecordList)
	if err != nil {
		log.Println("FindFinishScoreRecordListByQuestionId err ")
	}
	return err
}

func FindSelfScoreRecordByUserId(selfScoreRecord *[]ScoreRecord, examinerId int64) error {

	err := adapter.engine.Where("user_id=?", examinerId).Where("test_record_type =?", 0).Find(selfScoreRecord)
	if err != nil {
		log.Println("FindSelfScoreRecordByUserId err ")
	}
	return err
}

func GetTestScoreRecordByTestIdAndUserId(testScoreRecord *ScoreRecord, testId int64, examinerId string) error {

	_, err := adapter.engine.Where("user_id=?", examinerId).Where("test_id =?", testId).Where(" test_score_type !=?", 0).Get(testScoreRecord)
	if err != nil {
		log.Println("FindSelfScoreRecordByUserId err ")
	}
	return err
}

func CountTestByScore(question int64, score int64) (count int64, err1 error) {
	scoreRecord := new(ScoreRecord)
	count, err := adapter.engine.Where("score = ?", score).Where("test_record_type=1 or test_record_type=2  ").Where("question_id=?", question).Count(scoreRecord)
	if err != nil {
		log.Println("CountTestByScored err ")
	}
	return count, err
}

func (s *ScoreRecord) Update() error {
	code, err := adapter.engine.Where(builder.Eq{"test_id": s.TestId}).Update(s)
	if code == 0 || err != nil {
		log.Println("update ScoreRecord fail")
		log.Printf("%+v", err)
	}
	return err
}

func ListScoreRecordByUserId(userId int64) ([]ScoreRecord, error) {
	var records []ScoreRecord
	err := adapter.engine.Where(builder.Eq{"user_id": userId}).Find(&records)
	return records, err
}

func ListTeacherScoreByTestIds(userId int64, testIds []int64) ([]int64, error) {
	res := make([]int64, len(testIds))

	for i, testId := range testIds {
		var record ScoreRecord
		_, err := adapter.engine.Where(builder.Eq{"user_id": userId}).Where(builder.Eq{"test_id": testId}).Where("test_record_type=8").Get(&record) // 培训卷
		if err != nil {
			return nil, err
		}
		res[i] = record.Score
	}

	return res, nil
}

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
	Test_finish      int64      `json:"test_finish"`
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
func CountProblemFinishNumberByQuestionId(questionId int64)(count int64,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_record_type=?",6).Count(record)
	if err!=nil {
		log.Println("CountProblemFinishNumberByQuestionId err ")
	}
	return count,err1
}
func CountSelfScore(userId string,questionId int64)(count int64,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_record_type=0").Where("user_id=?", userId).Count(record)
	if err!=nil {
		log.Println("CountFailTestNumberByUserId err ")
	}
	return count ,err1
}
func CountFailTestNumberByUserId(userId string,questionId int64)(count int64,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_record_type=5").Where("user_id=?", userId).Count(record)
	if err!=nil {
		log.Println("CountFailTestNumberByUserId err ")
	}
	return count,err1
}
func CountProblemNumberByQuestionId(questionId int64)(count int64) {
	record :=new (ScoreRecord)
	count, err := x.Where("question_id = ?", questionId).Where("test_record_type=?",5).Count(record)
	if err!=nil {
		log.Println("CountProblemNumberByQuestionId err ")
	}
	return count
}
func CountArbitramentFinishNumberByQuestionId(questionId int64)(count int64,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_record_type=?",4).Count(record)
	if err!=nil {
		log.Println("CountArbitramentFinishNumberByQuestionId err ")
	}
	return count,err1
}



func CountFinishScoreNumberByQuestionId(questionId int64)(count int64 ,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_finish=1").Count(record)
	if err!=nil {
		log.Println("CountFinishScoreNumberByQuestionId err ")
	}
	return count,err1
}
func FindFinishTestNumberByUserId(scoreRecord *[]ScoreRecord,userId string,questionId int64)( err error) {
	err1 := x.Where("question_id = ?", questionId).Where("user_id=?",userId).Where("test_finish=1").Find(scoreRecord)
	if err!=nil {
		log.Println("FindFinishTestNumberByUserId err ")
	}
	return err1
}

func CountFirstScoreNumberByQuestionId(questionId int64)(count int64 ,err error) {
	record :=new (ScoreRecord)
	count, err1:= x.Where("question_id = ?", questionId).Where("test_record_type=?",1).Count(record)
	if err!=nil {
		log.Println("CountFirstScoreNumberByQuestionId err ")
	}
	return count ,err1
}
func CountSecondScoreNumberByQuestionId(questionId int64)(count int64 ,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_record_type=?",2).Count(record)
	if err!=nil {
		log.Println("CountSecondScoreNumberByQuestionId err ")
	}
	return count,err1
}
func CountThirdScoreNumberByQuestionId(questionId int64)(count int64,err error) {
	record :=new (ScoreRecord)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_record_type=?",3).Count(record)
	if err!=nil {
		log.Println("CountThirdScoreNumberByQuestionId err ")
	}
	return count,err1
}
func CountFinishTestNumberByUserId(userId string,questionId int64)(count int64,err1 error) {
	record :=new (ScoreRecord)
	count, err := x.Where("question_id = ?", questionId).Where("test_finish=1").Where("user_id=?", userId).Count(record)
	if err!=nil {
		log.Println("CountFinishTestNumberByUserId err ")
	}
	return count,err
}
func SumFinishScore(userId string,questionId int64)(sum float64,err1 error) {
	record :=new (ScoreRecord)
	sum, err := x.Where("question_id = ?", questionId).Where("test_record_type!=0").Where("user_id=?", userId).Sum(record,"score")
	if err!=nil {
		log.Println("SumFinishScore err ")
	}
	return sum,err
}
func FindFinishScoreRecordListByQuestionId (scoreRecordList *[]ScoreRecord , questionId int64) error{
	err := x.Where("question_id = ?",questionId).Find(scoreRecordList)
	if err!=nil {
		log.Println("FindFinishScoreRecordListByQuestionId err ")
	}
	return  err
}

func FindSelfScoreRecordByUserId(selfScoreRecord *[] ScoreRecord,examinerId string)error {

	err := x.Where("user_id=?", examinerId).Where("test_record_type =?",0).Find(selfScoreRecord)
	if err!=nil {
		log.Println("FindSelfScoreRecordByUserId err ")
	}
	return err

}
func GetTestScoreRecordByTestIdAndUserId(testScoreRecord *ScoreRecord,testId int64,examinerId string) error {

	_, err := x.Where("user_id=?", examinerId).Where("test_id =?", testId).Where(" test_score_type !=?", 0).Get(testScoreRecord)
	if err!=nil {
		log.Println("FindSelfScoreRecordByUserId err ")
	}
	return err
}
func CountTestByScore(question int64, score int64) (count int64,err1 error){
	scoreRecord := new(ScoreRecord)
	count, err := x.Where("score = ?", score).Where("question_id=?",question).Count(scoreRecord)
	if err!=nil {
		log.Println("CountTestByScored err ")
	}
	return count,err
}

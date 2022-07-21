package models

import (
	"log"

	"xorm.io/builder"
)

type UnderCorrectedPaper struct {
	UnderCorrected_id  int64  `json:"underCorrected_id" xorm:"pk autoincr"`
	User_id            string `json:"user_id"`
	Test_id            int64  `json:"test_id"`
	Question_id        int64  `json:"question_id"`
	Test_question_type int64  `json:"test_question_type"`
	Problem_type       int64  `json:"problem_type" xorm:"default(-1)"`
	Problem_message    string  `json:"test_message"`
	Self_score_id    string  `json:"self_score_id"`
}

func (u *UnderCorrectedPaper) GetUnderCorrectedPaper(userId string, testId int64) error {
	has, err := x.Where("test_id=?",testId).Where("user_id =?",userId).Get(u)
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
func (u *UnderCorrectedPaper) SupervisorDelete() error {
	code, err := x.Where(builder.Eq{"test_id": u.Test_id}).Where(" test_question_type =4 or  test_question_type =6 or  test_question_type =7").Delete(u)
	if code == 0 || err != nil {
		log.Println("delete fail")
	}
	return err
}
func (u *UnderCorrectedPaper) SelfMarkDelete() error {
	code, err := x.Where(builder.Eq{"test_id": u.Test_id}).Where(" test_question_type =0 ").Delete(u)
	if code == 0 || err != nil {
		log.Println("delete fail")
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

func GetDistributedPaperByUserId(id string, up *[]UnderCorrectedPaper) error {
	err := x.Where("user_id = ?", id).Find(up)
	if err != nil {
		log.Println("could not find any paper")
	}
	return err
}
func CountRemainingTestNumberByUserId(questionId int64 ,userId string)(count int64,err error) {
	underCorrectedPaper:=new (UnderCorrectedPaper)
	count, err1 := x.Where("question_id = ?", questionId).Where("user_id=?",userId).Where("test_question_type=1 or test_question_type=2").Count(underCorrectedPaper)
	if err!=nil {
		log.Println("CountRemainingTestNumberByUserId err ")
	}
	return count,err1
}
func CountArbitramentUnFinishNumberByQuestionId(questionId int64 )(count int64,err error) {
	underCorrectedPaper:=new (UnderCorrectedPaper)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_question_type=4").Count(underCorrectedPaper)
	if err!=nil {
		log.Println("CountRemainingTestNumberByUserId err ")
	}
	return count,err1
}
func CountProblemUnFinishNumberByQuestionId(questionId int64 )(count int64,err error) {
	underCorrectedPaper:=new (UnderCorrectedPaper)
	count, err1 := x.Where("question_id = ?", questionId).Where("test_question_type=6").Count(underCorrectedPaper)
	if err!=nil {
		log.Println("CountProblemUnFinishNumberByQuestionId err ")
	}
	return count,err1
}
func CountUnScoreTestNumberByQuestionId(questionId int64 )(count int64,err error) {
	underCorrectedPaper:=new (UnderCorrectedPaper)
	count, err1 := x.Where("question_id = ?", questionId).Count(underCorrectedPaper)
	if err!=nil {
		log.Println("CountUnScoreTestNumberByQuestionId err ")
	}
	return count,err1
}

func GetUnderCorrectedPaperByUserIdAndTestId(underCorrectedPaper * UnderCorrectedPaper ,userId string,testId int64) error {

	_, err := x.Where("user_id=?", userId).Where("test_id =?", testId).Where(" test_question_type !=?", 0).Get(underCorrectedPaper)
	if err!=nil {
		log.Println("GetUnderCorrectedPaperByUserIdAndTestId err ")
	}
 return err
}
func GetUnderCorrectedSupervisorPaperByTestQuestionTypeAndTestId(underCorrectedPaper * UnderCorrectedPaper ,testId int64) error {

	_, err := x.Where("test_id =?", testId).Where(" test_question_type =4 or  test_question_type =6 or  test_question_type =7").Get(underCorrectedPaper)
	if err!=nil {
		log.Println("GetUnderCorrectedPaperByUserIdAndTestId err ")
	}
 return err
}
func GetSelfScorePaperByTestQuestionTypeAndTestId(underCorrectedPaper * UnderCorrectedPaper ,testId int64,userId string) error {

	_, err := x.Where("test_id =?", testId).Where(" test_question_type =0").Where("user_Id = ?",userId).Get(underCorrectedPaper)
	if err!=nil {
		log.Println("GetSelfScorePaperByTestQuestionTypeAndTestId err ")
	}
 return err
}


func FindArbitramentUnderCorrectedPaperByQuestionId(arbitramentUnderCorrectedPaper *[] UnderCorrectedPaper,questionId int64)error{

	err := x.Where("question_id=?", questionId).Where(" test_question_type =?", 4).Find(arbitramentUnderCorrectedPaper)
	if err!=nil {
		log.Println("FindArbitramentUnderCorrectedPaperByQuestionId err ")
	}
   return err
}
func FindAllArbitramentUnderCorrectedPaper(arbitramentUnderCorrectedPaper *[] UnderCorrectedPaper,questionId int64)error{

	err := x.Where(" test_question_type =?", 4).Where("question_id= ?",questionId).Find(arbitramentUnderCorrectedPaper)
	if err!=nil {
		log.Println("FindAllArbitramentUnderCorrectedPaper err ")
	}
   return err
}

func FindProblemUnderCorrectedPaperByQuestionId(problemUnderCorrectedPaper *[] UnderCorrectedPaper,questionId int64) error{

	err := x.Where("question_id=?", questionId).Where(" test_question_type =?", 6).Find(problemUnderCorrectedPaper)
	if err!=nil {
		log.Println("FindProblemUnderCorrectedPaperByQuestionId err ")
	}
 return err
}
func FindSelfUnderCorrectedPaperByQuestionId(selfUnderCorrectedPaper *[] UnderCorrectedPaper,questionId int64) error{

	err := x.Where("question_id=?", questionId).Where(" test_question_type =?", 7).Find(selfUnderCorrectedPaper)
	if err!=nil {
		log.Println("FindSelfUnderCorrectedPaperByQuestionId err ")
	}
 return err
}
func FindSelfMarkPaperByQuestionId(selfMarkUnderCorrectedPaper *[]UnderCorrectedPaper,questionId int64) error{

	err := x.Where("question_id=?", questionId).Where(" test_question_type =?", 7).Find(selfMarkUnderCorrectedPaper)
	if err!=nil {
		log.Println("FindSelfMarkPaperByQuestionId err ")
	}
 return err
}
func FindProblemUnderCorrectedList(problemUnderCorrectedPaper *[] UnderCorrectedPaper) error{

	err := x.Where(" test_question_type =?", 6).Find(problemUnderCorrectedPaper)
	if err!=nil {
		log.Println("FindProblemUnderCorrectedList err ")
	}
 return err
}
func GetDistributedTestIdPaperByUserId(id string, up *[]int64) error {
	err := x.Table("under_corrected_paper").Select("test_id").Where("user_id = ?", id).Where(" test_question_type=0 or test_question_type=1 or test_question_type=2 or test_question_type=3").Find(up)
	if err != nil {
		log.Panic(err)
		log.Println("could not find any paper")
	}
	return err
}
func GetUnMarkSelfTestIdPaperByUserId(id string, up *[]int64) error {
	err := x.Table("under_corrected_paper").Select("test_id").Where("user_id = ?", id).Where(" test_question_type=0 ").Find(up)
	if err != nil {
		log.Panic(err)
		log.Println("GetUnMarkSelfTestIdPaperByUserId")
	}
	return err
}
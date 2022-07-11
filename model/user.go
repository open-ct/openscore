package model

import (
	"log"
	"math/rand"
	"time"

	"xorm.io/builder"
)

type User struct {
	UserId        int64     `json:"user_id" xorm:"pk autoincr"`
	UserName      string    `json:"user_name"`
	ExaminerCount string    `json:"examiner_count"`
	Password      string    `json:"password"`
	IdCard        string    `json:"id_card"`
	Address       string    `json:"address"`
	Tel           string    `json:"tel"`
	Email         string    `json:"email"`
	LoginTime     time.Time `json:"login_time"`
	ExistTime     time.Time `json:"exist_time"`
	OnlineTime    int64     `json:"online_time"`
	SubjectName   string    `json:"subject_name"`
	Status        int64     `json:"status"`
	UserType      int64     `json:"user_type"`
	IsDistribute  int64     `json:"is_distribute"`

	QuestionId int64 `json:"question_id"`
}

func (u *User) Insert() error {
	code, err := x.Insert(u)
	if code == 0 || err != nil {
		log.Println("insert user fail")
		log.Printf("%+v", err)
	}
	return err
}

func (u *User) GetUser(id int64) error {
	has, err := x.Where(builder.Eq{"user_id": id}).Get(u)
	if !has || err != nil {
		log.Println("could not found user")
	}
	return err
}
func (u *User) Update() error {
	_, err := x.Update(u)
	if err != nil {
		log.Println("could not Update user")
	}
	return err
}

func CountOnlineNumberUnDistribute() (count int64, err error) {
	user := new(User)
	count, err1 := x.Where("status = ? ", 1).Where(" user_type=?", 1).Where("is_distribute = ?", 0).Count(user)
	if err != nil {
		log.Println("CountOnlineNumber err ")
	}
	return count, err1
}
func CountOnlineUserNumberByQuestionId(questionId int64) (count int64, err error) {
	user := new(User)
	count, err1 := x.Where("status = ? ", 1).Where(" user_type=?", 1).Where("is_distribute = ?", 1).Where("question_id=?", questionId).Count(user)
	if err != nil {
		log.Println("CountOnlineNumber err ")
	}
	return count, err1
}
func FindOnlineUserNumberByQuestionId(users *[]User, questionId int64) error {
	err := x.Where("status = ? ", 1).Where(" user_type=?", 1).Where("is_distribute = ?", 1).Where("question_id=?", questionId).Find(users)
	if err != nil {
		log.Println("FindOnlineUserNumberByQuestionId err ")
	}
	return err
}

func FindUserNumberByQuestionId(users *[]User, questionId int64) error {
	err := x.Where(" user_type=?", 1).Where("is_distribute = ?", 1).Where("question_id=?", questionId).Find(users)
	if err != nil {
		log.Println("FindOnlineUserNumberByQuestionId err ")
	}
	return err
}

func FindUsers(u *[]User) error {
	err := x.Where("status = ? ", 1).Where(" user_type=?", 1).Where("is_distribute = ?", 0).Find(u)
	if err != nil {
		log.Println("could not FindUsers ")
		log.Println(err)
	}
	return err
}

func FindNewUserId(id1 int64, id2 int64, questionId int64) (newId int64) {
	var Ids []int64
	err := x.Table("user").Where("user_id !=?", id1).Where("user_id !=?", id2).Where("question_id=?", questionId).Where("status=?", 1).Select("user_id").Find(&Ids)
	if err != nil {
		log.Println("FindNewUserId err")
		log.Println(err)
	}
	k := rand.Intn(len(Ids) - 1)
	newId = Ids[k]
	return newId
}

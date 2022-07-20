package model

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"log"
	"math/rand"
	"xorm.io/builder"
)

type User struct {
	UserId         int64  `json:"user_id" xorm:"pk autoincr"`
	ExaminerCount  string `json:"examiner_count"`
	LoginTime      string `json:"login_time"`
	ExistTime      string `json:"exist_time"`
	OnlineTime     int64  `json:"online_time"`
	SubjectName    string `json:"subject_name"`
	IsOnlineStatus bool   `json:"is_online_status"`
	IsDistribute   bool   `json:"is_distribute"`
	CasdoorName    string `json:"casdoor_name" xorm:"not null unique"`
	QuestionId     int64  `json:"question_id"`
}

type UserInfo struct {
	User
	Casdoor *casdoorsdk.User
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

func (u *UserInfo) GetUserInfo(id int64) error {
	err := u.User.GetUser(id)
	if err != nil {
		return err
	}

	u.Casdoor, err = casdoorsdk.GetUser(u.User.CasdoorName)
	return err
}

func GetUserByCasdoorName(name string) (*User, error) {
	u := &User{}
	has, err := x.Where(builder.Eq{"casdoor_name": name}).Get(u)
	if !has {
		return nil, err
	}

	return u, err
}

func (u *User) UpdateCols(columns ...string) error {
	_, err := x.Cols(columns...).Update(u)
	if err != nil {
		log.Println("could not Update user")
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

func UpdateMemberOnlineStatus(userId int64, isOnline bool, time string) error {
	var user User
	if err := user.GetUser(userId); err != nil {
		return err
	}

	user.IsOnlineStatus = isOnline
	user.LoginTime = time

	return user.UpdateCols("is_online_status", "login_time")
}

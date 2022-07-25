package model

import (
	"log"
	"math/rand"
	"xorm.io/builder"
)

type User struct {
	UserId         int64  `json:"user_id" xorm:"pk autoincr"`
	ExaminerCount  string `json:"examiner_count"`
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	Address        string `json:"address"`
	Tel            string `json:"tel"`
	Email          string `json:"email"`
	IdCard         string `json:"id_card"`
	LoginTime      string `json:"login_time"`
	ExistTime      string `json:"exist_time"`
	OnlineTime     int64  `json:"online_time"`
	SubjectName    string `json:"subject_name"`
	IsOnlineStatus bool   `json:"is_online_status"`
	IsDistribute   bool   `json:"is_distribute"`
	QuestionId     int64  `json:"question_id"`
	UserType       int64  `json:"user_type"`
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

func GetUserByIdCard(idCard string) (*User, error) {
	u := &User{}
	has, err := x.Where(builder.Eq{"id_card": idCard}).Get(u)
	if !has || err != nil {
		log.Println("could not found user by IdCard")
		return nil, nil
	}
	return u, nil
}

func (u *User) GetUserInfo(id int64) error {
	// TODO
	return u.GetUser(id)
}

func (u *User) UpdateCols(columns ...string) error {
	_, err := x.ID(u.UserId).Cols(columns...).Update(u)
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

func FindUsers(u *[]User, subject string) error {
	err := x.Where("is_online_status = 1 AND user_type = 1 AND is_distribute = 0 AND subject_name = ?", subject).Find(u)
	if err != nil {
		log.Println("could not FindUsers ")
		log.Println(err)
	}
	return err
}

func CountOnlineUserNumberByQuestionId(questionId int64) (int64, error) {
	user := new(User)
	count, err := x.Where("is_online_status = 1").Where(" user_type = ?", 1).Where("is_distribute = ?", 1).Where("question_id=?", questionId).Count(user)
	if err != nil {
		log.Println("CountOnlineNumber err ")
	}
	return count, err
}

func FindOnlineUserNumberByQuestionId(users *[]User, questionId int64) error {
	return x.Where("is_online_status = ? ", 1).Where(" user_type=?", 1).Where("is_distribute = ?", 1).Where("question_id=?", questionId).Find(users)
}

func FindUserNumberByQuestionId(users *[]User, questionId int64) error {
	return x.Where(" user_type=?", 1).Where("is_distribute = ?", 1).Where("question_id=?", questionId).Find(users)
}

func FindNewUserId(id1 int64, id2 int64, questionId int64) (newId int64) {
	var Ids []int64
	err := x.Table("user").Where("user_id !=?", id1).Where("user_id !=?", id2).Where("question_id=?", questionId).Where("is_online_status=?", 1).Select("user_id").Find(&Ids)
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

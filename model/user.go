package model

import (
	"log"
	"math/rand"
	"time"

	"xorm.io/builder"
)

type User struct {
	User_id        string    `json:"user_id" xorm:"pk"`
	User_name      string    `json:"user_name"`
	Examiner_count string    `json:"examiner_count"`
	Password       string    `json:"password"`
	Id_card        string    `json:"id_card"`
	Address        string    `json:"address"`
	Tel            string    `json:"tel"`
	Email          string    `json:"email"`
	Login_time     time.Time `json:"login_time"`
	Exist_time     time.Time `json:"exist_time"`
	Online_time    int64     `json:"online_time"`
	Subject_name   string    `json:"subject_name"`
	Status         int64     `json:"status"`
	UserType       int64     `json:"userType"`
	IsDistribute   int64     `json:"isDistribute"`

	QuestionId int64 `json:"question_id"`
}

func (u *User) GetUser(id string) error {
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
func FindNewUserId(id1 string, id2 string, questionId int64) (newId string) {
	var Ids []string
	err := x.Table("user").Where("user_id !=?", id1).Where("user_id !=?", id2).Where("question_id=?", questionId).Where("status=?", 1).Select("user_id").Find(&Ids)
	if err != nil {
		log.Println("FindNewUserId err")
		log.Println(err)
	}
	k := rand.Intn(len(Ids) - 1)
	newId = Ids[k]
	return newId
}

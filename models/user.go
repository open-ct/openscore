package models

import (
	"log"
	"time"

	"xorm.io/builder"
)

type User struct {
	User_id        int64
	User_name      string
	Examiner_count string
	Password       string
	Id_card        string
	Address        string
	Tel            string
	Email          string
	Login_time     time.Time
	Exist_time     time.Time
	Online_time    float64
	Subject_name   string
	Status         int64
	UserType       int64
	IsDistribute   int64
}

func initUserModels() {
	err := x.Sync2(new(User))
	if err != nil {
		log.Println(err)
	}
}

func (u *User) GetUser(id int64) error {
	has, err := x.Where(builder.Eq{"user_id": id}).Get(u)
	if !has || err != nil {
		log.Println("could not found user")
	}
	return err
}

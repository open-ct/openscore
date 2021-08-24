package models

import (
	"log"
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
	Online_time    float64   `json:"online_time"`
	Subject_name   string    `json:"subject_name"`
	Status         int64     `json:"status"`
	UserType       int64     `json:"userType"`
	IsDistribute   int64     `json:"isDistribute"`
}

func (u *User) GetUser(id string) error {
	has, err := x.Where(builder.Eq{"user_id": id}).Get(u)
	if !has || err != nil {
		log.Println("could not found user")
	}
	return err
}

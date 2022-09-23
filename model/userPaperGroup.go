package model

import (
	"errors"
	"log"
	"xorm.io/builder"
)

type UserPaperGroup struct {
	Id      int64 `json:"id" xorm:"pk autoincr"`
	GroupId int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
}

func (u *UserPaperGroup) Update() error {
	code, err := adapter.engine.Where(builder.Eq{"id": u.Id}).Update(u)
	if code == 0 || err != nil {
		log.Println("update PaperGroup fail")
		log.Printf("%+v", err)
	}
	return err
}

func CreateUserPaperGroup(userId int64, groupId int64) error {
	data := &UserPaperGroup{
		GroupId: groupId,
		UserId:  userId,
	}

	_, err := adapter.engine.Insert(data)
	if err != nil {
		log.Println("CreateUserPaperGroup err ")
	}

	return err
}

func GetUserPaperGroupByUserId(id int64) (*UserPaperGroup, bool, error) {
	var group UserPaperGroup
	has, err := adapter.engine.Where("user_id=?", id).Get(&group)
	if err != nil {
		log.Println("GetGroupByGroupId err ")
		return nil, false, errors.New("GetGroupByGroupId")
	}

	return &group, has, nil
}

func ListUserPaperGroupByGroupId(id int64) ([]*UserPaperGroup, error) {
	var groups []*UserPaperGroup
	err := adapter.engine.Where(builder.Eq{"group_id": id}).Find(&groups)
	if err != nil {
		log.Println("ListUserPaperGroupByGroupId err ")
		return nil, err
	}

	return groups, nil
}

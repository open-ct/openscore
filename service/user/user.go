package user

import (
	"errors"
	"openscore/model"
	"openscore/pkg/token"
)

func Login(idCard string, pwd string) (string, error) {
	u, err := model.GetUserByIdCard(idCard)
	if err != nil {
		return "", err
	}
	if u == nil {
		return "", errors.New("user not found")
	}

	if u.Password != pwd {
		return "", errors.New("pwd not correct")
	}

	// 生成 auth token
	token, err := token.GenerateToken(&token.TokenPayload{
		Id:      u.UserId,
		Role:    u.UserType,
		Expired: token.GetExpiredTime(),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserInfo(token string) (*model.User, error) {
	user := &model.User{}
	err := user.GetUser(1)
	return user, err
}

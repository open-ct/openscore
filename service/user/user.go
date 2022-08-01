package user

import (
	"errors"
	"github.com/open-ct/openscore/model"
	"github.com/open-ct/openscore/pkg/token"
)

func Login(account string, pwd string) (string, int64, error) {
	u, err := model.GetUserByAccount(account)
	if err != nil {
		return "", 0, err
	}
	if u == nil {
		return "", 0, errors.New("user not found")
	}

	if u.Password != pwd {
		return "", 0, errors.New("pwd not correct")
	}

	// 生成 auth token
	token, err := token.GenerateToken(&token.TokenPayload{
		Id:      u.UserId,
		TypeId:  u.UserType,
		Expired: token.GetExpiredTime(),
	})
	if err != nil {
		return "", 0, err
	}
	//
	// if err := u.UpdateOnlineStatus(true, util.GetCurrentTime()); err != nil {
	// 	return token, 0, err
	// }

	return token, u.UserType, nil
}

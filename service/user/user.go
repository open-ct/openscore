package user

import (
	"errors"
	"github.com/open-ct/openscore/model"
	"github.com/open-ct/openscore/pkg/token"
	"github.com/open-ct/openscore/util"
)

func Login(account string, pwd string) (string, error) {
	u, err := model.GetUserByAccount(account)
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
		TypeId:  u.UserType,
		Expired: token.GetExpiredTime(),
	})
	if err != nil {
		return "", err
	}

	if err := u.UpdateOnlineStatus(true, util.GetCurrentTime()); err != nil {
		return token, err
	}

	return token, nil
}

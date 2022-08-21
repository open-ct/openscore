package user

import (
	"errors"

	"github.com/open-ct/openscore/model"
)

func Login(account string, pwd string) (int64, string, error) {
	u, err := model.GetUserByAccount(account)
	if err != nil {
		return 0, "", err
	}
	if u == nil {
		return 0, "", errors.New("user not found")
	}

	if u.Password != pwd {
		return 0, "", errors.New("pwd not correct")
	}

	// if err := u.UpdateOnlineStatus(true, util.GetCurrentTime()); err != nil {
	// 	return token, 0, err
	// }

	return u.UserId, u.UserType, nil
}

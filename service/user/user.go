package user

import (
	"errors"

	"github.com/open-ct/openscore/model"
)

func Login(account string, pwd string) (*model.User, error) {
	u, err := model.GetUserByAccount(account)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	if u.Password != pwd {
		return nil, errors.New("pwd not correct")
	}

	// if err := u.UpdateOnlineStatus(true, util.GetCurrentTime()); err != nil {
	// 	return token, 0, err
	// }

	return u, nil
}

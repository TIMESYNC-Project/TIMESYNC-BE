package service

import (
	"errors"
	"strings"
	"timesync-be/features/user"
	"timesync-be/helper"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}
func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	hashed, _ := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data already used"
		} else if strings.Contains(err.Error(), "empty") {
			msg = "email not allowed empty"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

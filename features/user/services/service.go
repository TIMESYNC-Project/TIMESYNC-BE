package service

import (
	"errors"
	"log"
	"strings"
	"timesync-be/config"
	"timesync-be/features/user"
	"timesync-be/helper"

	"github.com/golang-jwt/jwt"
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

func (uuc *userUseCase) Login(nip, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(nip)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "empty") {
			msg = "nip or password not allowed empty"
		} else {
			msg = "account not registered or server error"
		}
		return "", user.Core{}, errors.New(msg)
	}
	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))
	return useToken, res, nil

}

func (uuc *userUseCase) Delete(token interface{}) error {
	id := helper.ExtractToken(token)
	// if id <= 0 {
	// 	return errors.New("id not found, server error")
	// }
	err := uuc.qry.Delete(uint(id))
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

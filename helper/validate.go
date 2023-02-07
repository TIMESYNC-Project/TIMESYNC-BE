package helper

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"timesync-be/features/user"

	"github.com/go-playground/validator/v10"
)

func TypeFile(test multipart.File) bool {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	if fileType == "image/png" || fileType == "image/jpeg" {
		return true
	}
	return false
}

type UserValidate struct {
	Name        string `validate:"required,alphanum"`
	BirthOfDate string `validate:"required"`
	Email       string `validate:"required,email"`
	Gender      string `validate:"required,alpha"`
	Position    string `validate:"required"`
	Phone       string `validate:"required,numeric"`
	Address     string `validate:"required"`
	Password    string `validate:"required,min=3,alphanum"`
}

func CoreToRegVal(data user.Core) UserValidate {
	return UserValidate{
		Name:        data.Name,
		BirthOfDate: data.BirthOfDate,
		Email:       data.Email,
		Gender:      data.Gender,
		Position:    data.Position,
		Phone:       data.Phone,
		Address:     data.Address,
		Password:    data.Password,
	}
}

// func RegValToCore(data UserValidate) user.Core {
// 	return user.Core{
// 		Name:        data.Name,
// 		BirthOfDate: data.BirthOfDate,
// 		Email:       data.Email,
// 		Gender:      data.Gender,
// 		Position:    data.Position,
// 		Phone:       data.Phone,
// 		Address:     data.Address,
// 		Password:    data.Password,
// 	}
// }

func RegistrationValidate(data user.Core) error {
	validate := validator.New()
	val := CoreToRegVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := fmt.Sprintf("%s is %s", e.Field(), e.Tag())
			return errors.New(vlderror)
		}
	}

	return nil
}

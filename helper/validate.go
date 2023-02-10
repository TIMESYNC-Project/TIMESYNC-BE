package helper

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"timesync-be/features/approval"
	"timesync-be/features/user"

	"github.com/go-playground/validator/v10"
)

func TypeFile(test multipart.File) (string, error) {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	TipenamaFile := ""
	if fileType == "image/png" {
		TipenamaFile = ".png"
	} else {
		TipenamaFile = ".jpg"
	}
	if fileType == "image/png" || fileType == "image/jpeg" || fileType == "image/jpg" {
		return TipenamaFile, nil
	}
	return "", errors.New("file type not match")
}

func CsvTypeFile(fileHeader multipart.FileHeader) error {
	if fileHeader.Size == 0 {
		return errors.New("file empty")
	}
	src, err := fileHeader.Open()
	if err != nil {
		log.Println("open file error", err.Error())
		return errors.New("can't open file")
	}
	fileByte, _ := io.ReadAll(src)
	fileType := http.DetectContentType(fileByte)
	log.Println("|", fileType, "|")
	if fileType == "text/plain; charset=utf-8" {
		return nil
	}
	return errors.New("file type error, only csv can be upload")
}

type UserValidate struct {
	Name        string `validate:"required"`
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

type ApprovalValidate struct {
	Title       string `validate:"required"`
	StartDate   string `validate:"required"`
	EndDate     string `validate:"required"`
	Description string `validate:"required"`
}

func CoreToApprovalVal(data approval.Core) ApprovalValidate {
	return ApprovalValidate{
		Title:       data.Title,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		Description: data.Description,
	}
}

func ApprovalValidation(data approval.Core) error {
	validate := validator.New()
	val := CoreToApprovalVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := fmt.Sprintf("%s is %s", e.Field(), e.Tag())
			return errors.New(vlderror)
		}
	}

	return nil
}

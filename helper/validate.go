package helper

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"timesync-be/features/approval"
	"timesync-be/features/setting"
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
func RegistrationValidate(data user.Core) error {
	validate := validator.New()
	val := CoreToRegVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := ""
			if e.Field() == "Password" && e.Value() != "" {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
			if e.Value() == "" {
				vlderror = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
				return errors.New(vlderror)
			} else {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
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

func UpdateUserCheckValidation(data user.Core) error {
	validate := validator.New()
	if data.Password == "" && data.Phone == "" && data.Gender == "" && data.Email == "" {
		return nil
	}
	if data.Email != "" {
		err := validate.Var(data.Email, "email")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%s is not %s", data.Email, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Phone != "" {
		err := validate.Var(data.Phone, "numeric")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%s is not %s", data.Phone, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Password != "" {
		err := validate.Var(data.Password, "min=3,alphanum")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%s is not %s", data.Password, e.Tag())
			return errors.New(vlderror)
		}
	}

	return nil
}

func SettingValidate(data setting.Core) error {
	validate := validator.New()
	if data.Start == "" && data.End == "" && data.Tolerance == 0 && data.AnnualLeave == 0 {
		return nil
	}
	if len(data.Start) < 5 && len(data.Start) > 0 {
		return errors.New("wrong input format hour or minute, you should write like '22:30' or '02:09'")
	}
	if len(data.End) < 5 && len(data.End) > 0 {
		return errors.New("wrong input format hour or minute, you should write like '22:30' or '02:09'")
	}

	if len(data.Start) > 5 || len(data.End) > 5 {
		return errors.New("wrong input format hour or minute, you should write like '22:30' or '02:09'")
	}
	if data.Start != "" {
		m, _ := strconv.Atoi(data.Start[3:])
		h, _ := strconv.Atoi(data.Start[:2])
		if m > 59 || m < 0 {
			return errors.New("time start minute min 00 max 59")
		}
		if h >= 24 || h < 0 {
			return errors.New("time start hour min 00 max 23")
		}
		log.Println(h, m < 0)
		cekStart := strings.Replace(data.Start, ":", "", -1)
		err := validate.Var(cekStart, "numeric")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%s is not %s", data.Start, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.End != "" {
		m, _ := strconv.Atoi(data.End[3:])
		h, _ := strconv.Atoi(data.End[:2])
		if m > 59 || m < 0 {
			return errors.New("time end minute min 00 max 59")
		}
		if h >= 24 || h < 0 {
			return errors.New("time end hour min 00 max 23")
		}
		cekEnd := strings.Replace(data.End, ":", "", -1) // untuk memishkan char ":" <--
		log.Println(cekEnd)
		err := validate.Var(cekEnd, "numeric")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%s is not %s", data.End, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Tolerance != 0 {
		if data.Tolerance > 59 || data.Tolerance < 0 {
			return errors.New("tolerance minimum 0 minute and max 59 minute")
		}
		err := validate.Var(data.Tolerance, "numeric")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%d is not %s", data.Tolerance, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.AnnualLeave != 0 {
		if data.AnnualLeave < 0 {
			return errors.New("annual leave cannot less than 0")
		}
		err := validate.Var(data.AnnualLeave, "numeric")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("%d is not %s", data.AnnualLeave, e.Tag())
			return errors.New(vlderror)
		}
	}
	return nil
}

package helper

import (
	"log"
	"regexp"
	"timesync-be/features/user"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RegisterValidate struct {
	Name     string `validate:"required,alpha_space"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required,secure_password"`
}

type PasswordValidate struct {
	Password string `validate:"secure_password"`
}

type EmailValidate struct {
	Email string `validate:"email"`
}

type PhoneValidate struct {
	PhoneNumber string `validate:"numeric"`
}

func Validate(option string, data interface{}) interface{} {
	switch option {
	case "register":
		res := RegisterValidate{}
		if v, ok := data.(user.Core); ok {
			res.Name = v.Name
			res.Email = v.Email
			res.Phone = v.Password
			res.Password = v.Password
		}
		return res
	case "password":
		res := PasswordValidate{}
		if v, ok := data.(user.Core); ok {
			res.Password = v.Password
		}
		return res
	case "email":
		res := EmailValidate{}
		if v, ok := data.(user.Core); ok {
			res.Email = v.Email
		}
		return res
	case "phone":
		res := PhoneValidate{}
		if v, ok := data.(user.Core); ok {
			res.PhoneNumber = v.Phone
		}
		return res
	default:
		return nil
	}
}

func alphaSpace(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^[a-zA-Z\\s]+$", fl.Field().String())
	return match
}

func securePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false
	}
	if regexp.MustCompile(`^(?i)(password|1234|qwerty)`).MatchString(password) {
		return false
	}
	return true
}

func Validation(data interface{}) error {
	validate = validator.New()
	validate.RegisterValidation("alpha_space", alphaSpace)
	validate.RegisterValidation("secure_password", securePassword)
	err := validate.Struct(data)
	if err != nil {
		log.Println("log on helper validasi: ", err)
		return err
	}
	return nil
}

package services

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"
	"timesync-be/config"
	"timesync-be/features/user"
	"timesync-be/helper"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}

func (uuc *userUseCase) Register(token interface{}, newUser user.Core) (user.Core, error) {
	adminID := helper.ExtractToken(token)
	if len(newUser.Password) != 0 {
		//validation
		err := helper.RegistrationValidate(newUser)
		if err != nil {
			return user.Core{}, errors.New("validate: " + err.Error())
		}
	}
	hashed, _ := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)

	res, err := uuc.qry.Register(uint(adminID), newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data already used"
		} else if strings.Contains(err.Error(), "access denied") {
			msg = "access denied"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Login(nip, password string) (string, string, user.Core, error) {
	res, err := uuc.qry.Login(nip)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "empty") {
			msg = "nip or password not allowed empty"
		} else {
			msg = "account not registered or server error"
		}
		return "", "", user.Core{}, errors.New(msg)
	}
	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", "", user.Core{}, errors.New("password not matched")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	// claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))
	expired := helper.ExpiredToken()
	log.Println("waktu : ", time.Now())
	return expired, useToken, res, nil

}

func (uuc *userUseCase) Delete(token interface{}, employeeID uint) error {
	id := helper.ExtractToken(token)
	err := uuc.qry.Delete(uint(id), employeeID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

// Update implements user.UserService
func (uuc *userUseCase) Update(token interface{}, fileData multipart.FileHeader, updateData user.Core) (user.Core, error) {
	employeeID := helper.ExtractToken(token)
	if updateData.Password != "" {
		hashed, _ := helper.GeneratePassword(updateData.Password)
		updateData.Password = string(hashed)
	}
	if fileData.Size != 0 {
		if fileData.Size > 500000 {
			return user.Core{}, errors.New("size error")
		}

		fileName := uuid.NewV4().String()
		fileData.Filename = fileName + fileData.Filename[(len(fileData.Filename)-5):len(fileData.Filename)]
		src, err := fileData.Open()
		if err != nil {
			return user.Core{}, errors.New("error open fileData")
		}
		// Validasi Type
		if !helper.TypeFile(src) {
			return user.Core{}, errors.New("file type error only jpg or png file can be upload")
		}
		defer src.Close()
		uploadURL, err := helper.UploadToS3(fileData.Filename, src)
		if err != nil {
			return user.Core{}, errors.New("cannot upload to s3 server error")
		}
		updateData.ProfilePicture = uploadURL
	}
	res, err := uuc.qry.Update(uint(employeeID), updateData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email duplicated"
		} else if strings.Contains(err.Error(), "cannot modifed admin data") {
			msg = "cannot modifed admin data"
		} else {
			msg = "account not registered"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

// Csv implements user.UserService
func (uuc *userUseCase) Csv(fileData multipart.FileHeader) error {
	result := []user.Core{}
	if fileData.Filename != "" {
		src, err := fileData.Open()
		if err != nil {
			log.Println("open file error", err.Error())
			return errors.New("can't open file")
		}
		result = helper.ConvertCSV(src)
	} else {
		log.Panic(result)
	}
	err := uuc.qry.Csv(result)
	if err != nil {
		log.Println("query error")
		return errors.New("server error")
	}

	return nil

}

// Profile implements user.UserService
func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := uuc.qry.Profile(uint(userID))
	if err != nil {
		log.Println("data not found")
		return user.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// ProfileEmployee implements user.UserService
func (uuc *userUseCase) ProfileEmployee(userID uint) (user.Core, error) {
	res, err := uuc.qry.Profile(uint(userID))
	if err != nil {
		log.Println("data not found")
		return user.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// AdminEditEmployee implements user.UserService
func (uuc *userUseCase) AdminEditEmployee(employeeID uint, fileData multipart.FileHeader, updateData user.Core) (user.Core, error) {
	if updateData.Password != "" {
		hashed, _ := helper.GeneratePassword(updateData.Password)
		updateData.Password = string(hashed)
	}
	if fileData.Size != 0 {
		if fileData.Size > 500000 {
			return user.Core{}, errors.New("size error")
		}
		fileName := uuid.NewV4().String()
		fileData.Filename = fileName + fileData.Filename[(len(fileData.Filename)-5):len(fileData.Filename)]
		src, err := fileData.Open()
		if err != nil {
			return user.Core{}, errors.New("error open fileData")
		}
		defer src.Close()
		uploadURL, err := helper.UploadToS3(fileData.Filename, src)
		if err != nil {
			return user.Core{}, errors.New("cannot upload to s3 server error")
		}
		updateData.ProfilePicture = uploadURL
	}
	res, err := uuc.qry.Update(employeeID, updateData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email duplicated"
		} else if strings.Contains(err.Error(), "cannot modifed admin data") {
			msg = "cannot modifed admin data"
		} else {
			msg = "account not registered"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

// GetAllEmployee implements user.UserService
func (uuc *userUseCase) GetAllEmployee() ([]user.Core, error) {
	res, err := uuc.qry.GetAllEmployee()
	if err != nil {
		log.Println("data not found", err.Error())
		return []user.Core{}, errors.New("data not found")
	}
	return res, nil
}

// Search implements user.UserService
func (uuc *userUseCase) Search(token interface{}, quote string) ([]user.Core, error) {
	adminID := helper.ExtractToken(token)
	res, err := uuc.qry.Search(uint(adminID), quote)
	if err != nil {
		if strings.Contains(err.Error(), "access denied") {
			return []user.Core{}, errors.New("access denied")
		} else {
			return []user.Core{}, errors.New("data not found")
		}
	}
	return res, nil
}

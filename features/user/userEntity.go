package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	ProfilePicture string
	Name           string
	BirthOfDate    string
	Nip            string
	Email          string
	Gender         string
	Position       string
	Phone          string
	Address        string
	Password       string
	Role           string
	AnnualLeave    int
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Delete() echo.HandlerFunc
	// Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Csv() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(nip, password string) (string, Core, error)
	Delete(token interface{}) error
	// Profile(token interface{}) (interface{}, error)
	Update(employeeID uint, fileData multipart.FileHeader, updateData Core) (Core, error)
	Csv(fileHeader multipart.FileHeader) ([]Core, error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(nip string) (Core, error)
	// Profile(userID uint) (interface{}, error)
	Update(employeeID uint, updateData Core) (Core, error)
	Delete(id uint) error
	Csv(newUserBatch []Core) error
}

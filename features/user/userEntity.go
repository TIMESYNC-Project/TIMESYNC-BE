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
	Profile() echo.HandlerFunc
	ProfileEmployee() echo.HandlerFunc
	Update() echo.HandlerFunc
	Csv() echo.HandlerFunc
	AdminEditEmployee() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Csv(fileHeader multipart.FileHeader) ([]Core, error)
	Login(nip, password string) (string, Core, error)
	Delete(token interface{}, employeeID uint) error
	Profile(token interface{}) (Core, error)
	ProfileEmployee(userID uint) (Core, error)
	Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	AdminEditEmployee(employeeID uint, fileData multipart.FileHeader, updateData Core) (Core, error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(nip string) (Core, error)
	Profile(userID uint) (Core, error)
	Update(employeeID uint, updateData Core) (Core, error)
	Delete(adminID uint, employeeID uint) error
	Csv(newUserBatch []Core) error
}

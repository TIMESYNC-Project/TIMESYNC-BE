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
	Email          string `validate:"omitempty,email"`
	Gender         string
	Position       string
	Phone          string
	Address        string
	Password       string `validate:"min=8,omitempty"`
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
	GetAllEmployee() echo.HandlerFunc
	Search() echo.HandlerFunc
}

type UserService interface {
	Register(token interface{}, newUser Core) (Core, error)
	Csv(fileHeader multipart.FileHeader) error
	Login(nip, password string) (string, string, Core, error)
	Delete(token interface{}, employeeID uint) error
	Profile(token interface{}) (Core, error)
	ProfileEmployee(userID uint) (Core, error)
	Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	AdminEditEmployee(token interface{}, employeeID uint, fileData multipart.FileHeader, updateData Core) (Core, error)
	GetAllEmployee() ([]Core, error)
	Search(token interface{}, quote string) ([]Core, error)
}

type UserData interface {
	Register(adminID uint, newUser Core) (Core, error)
	Login(nip string) (Core, error)
	Profile(userID uint) (Core, error)
	Update(employeeID uint, updateData Core) (Core, error)
	UpdateByAdmin(adminID uint, employeeID uint, updateData Core) (Core, error)
	Delete(adminID uint, employeeID uint) error
	Csv(newUserBatch []Core) error
	GetAllEmployee() ([]Core, error)
	Search(adminID uint, quote string) ([]Core, error)
}

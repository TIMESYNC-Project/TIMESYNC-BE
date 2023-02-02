package user

import "github.com/labstack/echo"

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
	// Profile() echo.HandlerFunc
	// Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(nip, password string) (string, Core, error)
	// Profile(token interface{}) (interface{}, error)
	// Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	Delete(token interface{}) error
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(nip string) (Core, error)
	// Profile(userID uint) (interface{}, error)
	// Update(id uint, updateData Core) (Core, error)
	Delete(id uint) error
}

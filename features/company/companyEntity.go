package company

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	Picture        string
	Name           string
	Email          string
	Description    string
	CompanyAddress string
	CompanyPhone   string
	SocialMedia    string
}

type CompanyHandler interface {
	GetProfile() echo.HandlerFunc
	EditProfile() echo.HandlerFunc
}

type CompanyService interface {
	GetProfile() (Core, error)
	EditProfile(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
}
type CompanyData interface {
	GetProfile() (Core, error)
	EditProfile(adminID uint, updateData Core) (Core, error)
}

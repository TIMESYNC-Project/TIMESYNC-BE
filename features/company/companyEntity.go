package company

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint   `json:"id"`
	Picture        string `json:"company_picture"`
	Name           string `json:"company_name"`
	Email          string `json:"company_email"`
	Description    string `json:"description"`
	CompanyAddress string `json:"company_address"`
	CompanyPhone   string `json:"company_phone"`
	SocialMedia    string `json:"sosmed"`
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

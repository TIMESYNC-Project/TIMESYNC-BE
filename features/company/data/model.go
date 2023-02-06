package data

import (
	"timesync-be/features/company"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Picture        string
	Name           string
	Email          string
	SocialMedia    string
	Description    string
	CompanyPhone   string
	CompanyAddress string
}

func DataToCore(data Company) company.Core {
	return company.Core{
		ID:             data.ID,
		Picture:        data.Picture,
		Name:           data.Name,
		Email:          data.Email,
		SocialMedia:    data.SocialMedia,
		Description:    data.Description,
		CompanyPhone:   data.CompanyPhone,
		CompanyAddress: data.CompanyAddress,
	}
}

func CoreToData(data company.Core) Company {
	return Company{
		Model:          gorm.Model{ID: data.ID},
		Picture:        data.Picture,
		Name:           data.Name,
		Email:          data.Email,
		SocialMedia:    data.SocialMedia,
		Description:    data.Description,
		CompanyPhone:   data.CompanyPhone,
		CompanyAddress: data.CompanyAddress,
	}
}

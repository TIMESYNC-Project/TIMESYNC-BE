package handler

import (
	"mime/multipart"
	"timesync-be/features/company"
)

type EditRequest struct {
	Picture        string `json:"company_picture" form:"company_picture"`
	Name           string `json:"company_name" form:"company_name"`
	Email          string `json:"company_email" form:"company_email"`
	SocialMedia    string `json:"sosmed" form:"sosmed"`
	Description    string `json:"description" form:"description"`
	CompanyPhone   string `json:"company_phone" form:"company_phone"`
	CompanyAddress string `json:"company_address" form:"company_address"`
	FileHeader     multipart.FileHeader
}

func ReqToCore(data interface{}) *company.Core {
	res := company.Core{}
	switch data.(type) {
	case EditRequest:
		cnv := data.(EditRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.SocialMedia = cnv.SocialMedia
		res.Description = cnv.Description
		res.CompanyPhone = cnv.CompanyPhone
		res.CompanyAddress = cnv.CompanyAddress
	default:
		return nil
	}
	return &res
}

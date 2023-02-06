package handler

import (
	"mime/multipart"
	"timesync-be/features/company"
)

type EditRequest struct {
	Picture        string `json:"picture" form:"picture"`
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	SocialMedia    string `json:"social_media" form:"social_media"`
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
		res.Picture = cnv.Picture
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

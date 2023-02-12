package services

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"timesync-be/features/company"
	"timesync-be/helper"
)

type companyUseCase struct {
	qry company.CompanyData
}

func New(uc company.CompanyData) company.CompanyService {
	return &companyUseCase{
		qry: uc,
	}
}

// GetProfile implements company.CompanyService
func (cuc *companyUseCase) GetProfile() (company.Core, error) {
	res, err := cuc.qry.GetProfile()
	if err != nil {
		log.Println("data not found")
		return company.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// EditProfile implements company.CompanyService
func (cuc *companyUseCase) EditProfile(token interface{}, fileData multipart.FileHeader, updateData company.Core) (company.Core, error) {
	adminID := helper.ExtractToken(token)
	// kondisi dibawah dilakukan agar foto bisa kosong dan agar unit testing tidak error
	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return company.Core{}, errors.New("validate: " + err.Error())
	}
	updateData.Picture = url
	// if fileData.Size != 0 {
	// }
	res, err := cuc.qry.EditProfile(uint(adminID), updateData)
	if err != nil {
		log.Println("query error")
		if strings.Contains(err.Error(), "access") {
			return company.Core{}, errors.New("access denied")
		} else if strings.Contains(err.Error(), "no data") {
			return company.Core{}, errors.New("no data updated")
		} else {
			return company.Core{}, err
		}
	}
	return res, nil
}

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
	if fileData.Size != 0 {
		if fileData.Filename != "" {
			res, err := helper.GetUrlImagesFromAWS(fileData)
			if err != nil {
				return company.Core{}, err
			}
			updateData.Picture = res
		}
	}
	res, err := cuc.qry.EditProfile(uint(adminID), updateData)
	if err != nil {
		log.Println("query error")
		if strings.Contains(err.Error(), "access") {
			return company.Core{}, errors.New("access denied")
		} else {
			return company.Core{}, errors.New("query error, problem with server")
		}
	}
	return res, nil
}

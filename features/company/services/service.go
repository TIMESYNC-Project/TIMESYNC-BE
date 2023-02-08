package services

import (
	"errors"
	"log"
	"mime/multipart"
	"timesync-be/features/company"
	"timesync-be/helper"

	uuid "github.com/satori/go.uuid"
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
	if fileData.Filename != "" {
		if fileData.Size > 500000 {
			return company.Core{}, errors.New("size error")
		}
		file, err := fileData.Open()
		if err != nil {
			return company.Core{}, errors.New("error open fileData")
		}
		_, err = helper.TypeFile(file)
		if err != nil {
			return company.Core{}, errors.New("error open fileData")
		}
		fileName := uuid.NewV4().String()
		fileData.Filename = fileName + fileData.Filename[(len(fileData.Filename)-5):len(fileData.Filename)]
		src, err := fileData.Open()
		if err != nil {
			return company.Core{}, errors.New("error open fileData")
		}
		defer src.Close()
		uploadURL, err := helper.UploadToS3(fileData.Filename, src)
		if err != nil {
			return company.Core{}, errors.New("cannot upload to s3 server error")
		}
		updateData.Picture = uploadURL
	}
	res, err := cuc.qry.EditProfile(uint(adminID), updateData)
	if err != nil {
		log.Println("data not found")
		return company.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

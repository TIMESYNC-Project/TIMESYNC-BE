package services

import (
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
	"timesync-be/features/company"
	"timesync-be/helper"
	"timesync-be/mocks"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	repo := mocks.NewCompanyData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	resData := company.Core{
		ID:             1,
		Picture:        imageTrueCnv.Filename,
		Name:           "Timesync Company",
		Email:          "timesync@company.co.id",
		Description:    "IT Company",
		CompanyAddress: "Jl. Jalandikuburan, no.3, Jakarta Selatan",
		CompanyPhone:   "080898",
		SocialMedia:    "@timesync",
	}

	t.Run("success get company profile", func(t *testing.T) {
		repo.On("GetProfile").Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetProfile()
		assert.Nil(t, err)
		assert.Equal(t, res.ID, resData.ID)
		repo.AssertExpectations(t)
	})
	t.Run("server problem", func(t *testing.T) {
		repo.On("GetProfile").Return(company.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		res, err := srv.GetProfile()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, int(res.ID))
		repo.AssertExpectations(t)
	})
}

func TestEditProfile(t *testing.T) {
	repo := mocks.NewCompanyData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	inputData := company.Core{
		ID:             1,
		Name:           "Timesync Company",
		Email:          "timesync@company.co.id",
		Description:    "IT Company",
		CompanyAddress: "Jl. Jalandikuburan, no.3, Jakarta Selatan",
		CompanyPhone:   "080898",
		SocialMedia:    "@timesync",
	}
	resData := company.Core{
		ID:             1,
		Name:           "Timesync Company",
		Email:          "timesync@company.co.id",
		Description:    "IT Company",
		CompanyAddress: "Jl. Jalandikuburan, no.3, Jakarta Selatan",
		CompanyPhone:   "080898",
		SocialMedia:    "@timesync",
	}
	t.Run("success update company profile", func(t *testing.T) {
		repo.On("EditProfile", uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.EditProfile(pToken, *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Name, res.Name)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("EditProfile", uint(1), inputData).Return(company.Core{}, errors.New("query error, problem with server"))
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.EditProfile(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("EditProfile", uint(2), inputData).Return(company.Core{}, errors.New("access denied"))
		srv := New(repo)

		_, token := helper.GenerateToken(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.EditProfile(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "denied")
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "TimeSyncUnitTesting.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("dari file header unit testing approval", err.Error())
		}
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.EditProfile(pToken, *headerFake, inputData)
		assert.ErrorContains(t, err, "type")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("EditProfile", uint(2), company.Core{}).Return(company.Core{}, errors.New("no data updated"))
		srv := New(repo)

		_, token := helper.GenerateToken(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.EditProfile(pToken, *imageTrueCnv, company.Core{})
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "no data updated")
		repo.AssertExpectations(t)
	})
	// t.Run("invalid file validation", func(t *testing.T) {
	// 	filePathFake := filepath.Join("..", "..", "..", "Samurai Ghost Mask.png")
	// 	headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
	// 	if err != nil {
	// 		log.Panic("dari file header unit testing approval", err.Error())
	// 	}
	// 	srv := New(repo)
	// 	_, token := helper.GenerateToken(1)
	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true
	// 	res, err := srv.EditProfile(pToken, *headerFake, inputData)
	// 	assert.ErrorContains(t, err, "size")
	// 	assert.Error(t, err)
	// 	assert.Equal(t, uint(0), res.ID)
	// 	repo.AssertExpectations(t)

	// })
}

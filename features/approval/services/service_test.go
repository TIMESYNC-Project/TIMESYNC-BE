package services

import (
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
	"timesync-be/features/approval"
	"timesync-be/helper"
	"timesync-be/mocks"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostApproval(t *testing.T) {
	repo := mocks.NewApprovalData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	inputData := approval.Core{
		ID:          0,
		Title:       "Sakit",
		StartDate:   "2023-02-01",
		EndDate:     "2023-02-04",
		Description: "maaf pak tidak bisa hadir karena sakit",
		Status:      "pending",
	}
	resData := approval.Core{
		ID:          1,
		Title:       "Sakit",
		StartDate:   "2023-02-01",
		EndDate:     "2023-02-04",
		Description: "maaf pak tidak bisa hadir karena sakit",
		Status:      "pending",
	}

	t.Run("success post approval", func(t *testing.T) {
		repo.On("PostApproval", uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.PostApproval(pToken, *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("PostApproval", uint(1), mock.Anything).Return(approval.Core{}, errors.New("server error, failed to query")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.PostApproval(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, approval.Core{}, res)
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
		res, err := srv.PostApproval(pToken, *headerFake, inputData)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("invalid input validation", func(t *testing.T) {
		inputDataFake := approval.Core{
			Title:       "",
			StartDate:   "2023-02-01",
			EndDate:     "2023-02-04",
			Description: "",
		}
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.PostApproval(pToken, *imageTrueCnv, inputDataFake)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})
}

func TestGetApproval(t *testing.T) {
	repo := mocks.NewApprovalData(t)
	srv := New(repo)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	resData := []approval.Core{{
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: imageTrueCnv.Filename,
		Status:        "pending",
	}}

	t.Run("success get approval record", func(t *testing.T) {
		repo.On("GetApproval").Return(resData, nil).Once()
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetApproval()
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetApproval").Return([]approval.Core{}, errors.New("data not found")).Once()
		res, err := srv.GetApproval()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetApproval").Return([]approval.Core{}, errors.New("server problem")).Once()
		res, err := srv.GetApproval()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}

func TestApprovalDetail(t *testing.T) {
	repo := mocks.NewApprovalData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	resData := approval.Core{
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: imageTrueCnv.Filename,
		Status:        "pending",
	}

	t.Run("success get approval detail", func(t *testing.T) {
		repo.On("ApprovalDetail", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.ApprovalDetail(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("data not found", func(t *testing.T) {
		repo.On("ApprovalDetail", uint(1)).Return(approval.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		res, err := srv.ApprovalDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("server problem", func(t *testing.T) {
		repo.On("ApprovalDetail", uint(1)).Return(approval.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		res, err := srv.ApprovalDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestEmployeeApprovalRecord(t *testing.T) {
	repo := mocks.NewApprovalData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	resData := []approval.Core{{
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: imageTrueCnv.Filename,
		Status:        "pending",
	}}

	t.Run("success get employee approval record", func(t *testing.T) {
		repo.On("EmployeeApprovalRecord", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.EmployeeApprovalRecord(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("EmployeeApprovalRecord", uint(1)).Return([]approval.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.EmployeeApprovalRecord(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("EmployeeApprovalRecord", uint(1)).Return([]approval.Core{}, errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.EmployeeApprovalRecord(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}

func TestUpdateApproval(t *testing.T) {
	repo := mocks.NewApprovalData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	inputData := approval.Core{
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: "ERD.png",
		Status:        "pending",
	}
	resData := approval.Core{
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: imageTrueCnv.Filename,
		Status:        "approved",
	}

	t.Run("success update approval", func(t *testing.T) {
		repo.On("UpdateApproval", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateApproval(pToken, uint(1), inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Title, res.Title)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("UpdateApproval", uint(1), uint(1), inputData).Return(approval.Core{}, errors.New("data not found")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateApproval(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("unauthorized request", func(t *testing.T) {
		repo.On("UpdateApproval", uint(1), uint(1), inputData).Return(approval.Core{}, errors.New("unauthorized request")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateApproval(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "unauthorized")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("UpdateApproval", uint(1), uint(1), inputData).Return(approval.Core{}, errors.New("server problem"))
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateApproval(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "unable")
		repo.AssertExpectations(t)
	})
}

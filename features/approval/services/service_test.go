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
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: "ERD.png",
	}
	resData := approval.Core{
		ID:            1,
		Title:         "Sakit",
		StartDate:     "2023-02-01",
		EndDate:       "2023-02-04",
		Description:   "maaf pak tidak bisa hadir karena sakit",
		ApprovalImage: imageTrueCnv.Filename,
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
		assert.Equal(t, resData.Title, res.Title)
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt token", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.PostApproval(pToken, *imageTrueCnv, inputData)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("PostApproval", uint(1), inputData).Return(approval.Core{}, errors.New("data not found")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.PostApproval(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("PostApproval", uint(1), inputData).Return(approval.Core{}, errors.New("server problem"))
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.PostApproval(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

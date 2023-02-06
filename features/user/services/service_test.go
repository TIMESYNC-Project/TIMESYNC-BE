package services

import (
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
	"timesync-be/features/user"
	"timesync-be/helper"
	"timesync-be/mocks"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputData := user.Core{Name: "Fauzi", Email: "fauzi@example.com", Phone: "08123", Password: "123"}
	resData := user.Core{ID: uint(1), Name: "Fauzi", Email: "fauzi@example.com", Phone: "08123"}

	t.Run("success creating account", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Phone, res.Phone)
		repo.AssertExpectations(t)
	})

	t.Run("all input must fill", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("email or password not allowed empty")).Once()
		srv := New(repo)
		res, err := srv.Register(user.Core{})
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not allowed empty")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})
	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "used")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputEmail := "fauzi@gmail.com"
	hashed, _ := helper.GeneratePassword("123")
	resData := user.Core{ID: uint(1), Name: "Fauzi", Email: "fauzi@gmail.com", Password: hashed}
	t.Run("login success", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Name, res.Name)
		repo.AssertExpectations(t)
	})
	t.Run("account not found", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "123")
		assert.NotNil(t, token)
		assert.ErrorContains(t, err, "not")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "hitler@example.com"
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		_, res, err := srv.Login(inputEmail, "342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "hitler@example.com"
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("nip or password not allowed empty")).Once()
		srv := New(repo)
		_, res, err := srv.Login(inputEmail, "")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "empty")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := user.Core{ID: uint(1), Name: "fauzi", Email: "fauzi@example.com", Phone: "08123"}
	t.Run("success show profile", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData, res)
		repo.AssertExpectations(t)
	})
	t.Run("account not found", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("query error, problem with server")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestProfileEmployee(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := user.Core{ID: uint(1), Name: "fauzi", Email: "fauzi@example.com", Phone: "08123"}
	t.Run("success show profile", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.ProfileEmployee(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData, res)
		repo.AssertExpectations(t)
	})
	t.Run("account not found", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("query error, problem with server")).Once()
		srv := New(repo)

		res, err := srv.ProfileEmployee(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	// imageFalse, _ := os.Open(filePath)
	// imageFalseCnv := &multipart.FileHeader{
	// 	Filename: imageFalse.Name(),
	// }
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	inputData := user.Core{ID: 1, Name: "Alif", Phone: "08123", ProfilePicture: "ERD.png"}
	resData := user.Core{ID: 1, Name: "Alif", Phone: "08123", ProfilePicture: imageTrueCnv.Filename}
	t.Run("success updating account", func(t *testing.T) {
		repo.On("Update", uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("fail updating account", func(t *testing.T) {
		repo.On("Update", uint(1), inputData).Return(user.Core{}, errors.New("query error,update fail")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("deleting account successful", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(2)).Return(nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken, uint(2))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	// internal server error, account fail to delete
	t.Run("internal server error, account fail to delete", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(2)).Return(errors.New("no user has delete")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, uint(2))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		repo.AssertExpectations(t)
	})
}

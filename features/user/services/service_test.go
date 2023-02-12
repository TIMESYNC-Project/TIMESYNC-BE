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
	inputData := user.Core{Name: "Fauzi", Email: "fauzi@example.com", Phone: "08123"}
	resData := user.Core{ID: uint(1), Name: "Fauzi", Email: "fauzi@example.com", Phone: "08123"}

	t.Run("success creating account", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Phone, res.Phone)
		repo.AssertExpectations(t)
	})
	t.Run("invalid validation", func(t *testing.T) {
		inputDataSalah := user.Core{Name: "Fauzi%&*", Password: "123a#$%"}
		// repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputDataSalah)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})
	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "email already registered")
		repo.AssertExpectations(t)
	})

	t.Run("access denied", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputEmail := "fauzi@gmail.com"
	passwordHashed := helper.GeneratePassword("123")
	resData := user.Core{ID: uint(1), Name: "Fauzi", Email: "fauzi@gmail.com", Password: passwordHashed}
	t.Run("login success", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		_, token, res, err := srv.Login(inputEmail, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Name, res.Name)
		repo.AssertExpectations(t)
	})
	t.Run("account not found", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		_, token, res, err := srv.Login(inputEmail, "123")
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
		_, _, res, err := srv.Login(inputEmail, "342")
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
		_, _, res, err := srv.Login(inputEmail, "")
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
	//mock file kosong agar tidak terupdate di server aws
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	// log.Panic(imageTrueCnv.Header)
	inputData := user.Core{ID: 1, Name: "Alif", Phone: "08123"}
	resData := user.Core{ID: 1, Name: "Alif", Phone: "08123"}

	t.Run("success updating account", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(resData, nil).Once()
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
		repo.On("Update", uint(1), mock.Anything).Return(user.Core{}, errors.New("user not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not registered")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("email duplicated", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(user.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "TimeSyncUnitTesting.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("dari file header", err.Error())
		}
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *headerFake, inputData)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("invalid input validation", func(t *testing.T) {
		inputDataFake := user.Core{Password: "08123##"}
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, *imageTrueCnv, inputDataFake)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})
}

func TestAdminEditEmploye(t *testing.T) {
	repo := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
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
		repo.On("UpdateByAdmin", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("UpdateByAdmin", uint(1), uint(1), inputData).Return(user.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("fail updating account", func(t *testing.T) {
		repo.On("UpdateByAdmin", uint(1), uint(1), inputData).Return(user.Core{}, errors.New("user not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "account not registered")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("email duplicated", func(t *testing.T) {
		repo.On("UpdateByAdmin", uint(1), uint(1), inputData).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "TimeSyncUnitTesting.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("dari file header", err.Error())
		}
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *headerFake, inputData)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("invalid input validation", func(t *testing.T) {
		inputDataFake := user.Core{Name: "Alif", Phone: "08123##"}
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *imageTrueCnv, inputDataFake)
		assert.ErrorContains(t, err, "validate")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("no data updated", func(t *testing.T) {
		repo.On("UpdateByAdmin", uint(1), uint(1), inputData).Return(user.Core{}, errors.New("no data updated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.AdminEditEmployee(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Error(t, err)
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

func TestGetAllEmployee(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := []user.Core{
		{
			ID:    1,
			Name:  "Fauzi",
			Email: "fauzi@gmail.com",
			Phone: "081234",
		},
	}
	t.Run("get all employee successful", func(t *testing.T) {
		repo.On("GetAllEmployee").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetAllEmployee()
		assert.Equal(t, res[0].ID, resData[0].ID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	// internal server error, account fail to GetAllEmployee
	t.Run("internal server error, account fail to GetAllEmployee", func(t *testing.T) {
		repo.On("GetAllEmployee").Return([]user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		res, err := srv.GetAllEmployee()
		assert.Equal(t, res, []user.Core{})
		assert.ErrorContains(t, err, "data not found")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestSearch(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := []user.Core{{ID: 1, Name: "Fauzi", Nip: "23002"}}
	t.Run("success Found", func(t *testing.T) {
		repo.On("Search", uint(1), "eko").Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Search(pToken, "eko")
		assert.Nil(t, err)
		assert.Equal(t, resData[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})
	t.Run("Not found", func(t *testing.T) {
		repo.On("Search", uint(1), "123ad").Return([]user.Core{}, errors.New("no user found")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Search(pToken, "123ad")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data")
		assert.Equal(t, []user.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("Search", uint(2), "123ad").Return([]user.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Search(pToken, "123ad")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, []user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestCsv(t *testing.T) {
	repo := mocks.NewUserData(t)
	srv := New(repo)

	inputData := []user.Core{
		{
			ID:          1,
			Name:        "Fauzi Sofyan",
			BirthOfDate: "2000-01-31",
			Email:       "fauzilax@gmail.com",
			Gender:      "male",
			Position:    "backendengineer",
			Phone:       "0813456",
			Address:     "bandung",
		},
	}
	filePath := filepath.Join("..", "..", "..", "TimeSyncUnitTesting.csv")
	header, err := helper.UnitTestingUploadFileMock(filePath)
	if err != nil {
		log.Panic("dari file header", err.Error())
	}

	if header.Filename != "" {
		inputData = helper.ConvertCSV(*header)
	}
	// MULAI TESTING
	t.Run("success creating account from csv", func(t *testing.T) {
		repo.On("Csv", inputData).Return(nil).Once()
		err = srv.Csv(*header)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Duplicate", func(t *testing.T) {
		repo.On("Csv", inputData).Return(errors.New("Duplicate")).Once()
		err = srv.Csv(*header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "some email already registered in data entry")
		repo.AssertExpectations(t)
	})
	t.Run("Server error", func(t *testing.T) {
		repo.On("Csv", inputData).Return(errors.New("server error")).Once()
		err = srv.Csv(*header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		repo.AssertExpectations(t)
	})
	t.Run("invalid validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "Unit_Testing.PNG")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("dari file header", err.Error())
		}
		// repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err = srv.Csv(*headerFake)
		assert.ErrorContains(t, err, "validate")
		repo.AssertExpectations(t)

	})
}

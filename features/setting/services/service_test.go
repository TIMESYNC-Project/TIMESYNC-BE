package services

import (
	"errors"
	"testing"
	"timesync-be/features/setting"
	"timesync-be/helper"
	"timesync-be/mocks"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestEditSetting(t *testing.T) {
	data := mocks.NewSettingData(t)
	inputData := setting.Core{ID: 1, Start: "08:30", End: "16:30", Tolerance: 15, AnnualLeave: 14}
	resData := setting.Core{ID: 1, Start: "08:30", End: "16:30", Tolerance: 15, AnnualLeave: 14}

	t.Run("success update setting", func(t *testing.T) {
		data.On("EditSetting", uint(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.EditSetting(mockToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Start, res.Start)
		data.AssertExpectations(t)
	})

	t.Run("all input must fill", func(t *testing.T) {
		data.On("EditSetting", uint(1), inputData).Return(setting.Core{}, errors.New("update fail")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.EditSetting(mockToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "query error, problem with server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

}

func TestGetSetting(t *testing.T) {
	data := mocks.NewSettingData(t)
	resData := setting.Core{ID: 1, Start: "08:30", End: "16:30", Tolerance: 15, AnnualLeave: 14}

	t.Run("success update setting", func(t *testing.T) {
		data.On("GetSetting").Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.GetSetting()
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Start, res.Start)
		data.AssertExpectations(t)
	})

	t.Run("all input must fill", func(t *testing.T) {
		data.On("GetSetting").Return(setting.Core{}, errors.New("query error, problem with server")).Once()
		srv := New(data)
		res, err := srv.GetSetting()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "query error, problem with server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

}

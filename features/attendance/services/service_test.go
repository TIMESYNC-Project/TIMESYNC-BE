package services

import (
	"errors"
	"testing"
	"timesync-be/features/attendance"
	"timesync-be/helper"
	"timesync-be/mocks"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestClockIn(t *testing.T) {
	data := mocks.NewAttendanceData(t)
	resData := attendance.Core{
		ClockIn:         "07:50",
		ClockInLocation: "Jalan Soekarno hatta Bandung Jawa Barat",
		ClockInOSM:      "https://www.openstreetmap.org/#map=16/-6.4096/106.8185",
	}

	t.Run("success clockin", func(t *testing.T) {
		data.On("ClockIn", uint(1), "-6.4096", "106.8185").Return(resData, nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.ClockIn(mockToken, "-6.4096", "106.8185")
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.ClockIn, res.ClockIn)
		data.AssertExpectations(t)
	})

	t.Run("server error", func(t *testing.T) {
		data.On("ClockIn", uint(1), "-6.4096", "106.8185").Return(attendance.Core{}, errors.New("server error, location not found")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.ClockIn(mockToken, "-6.4096", "106.8185")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("already clockin", func(t *testing.T) {
		data.On("ClockIn", uint(1), "-6.4096", "106.8185").Return(attendance.Core{}, errors.New("you already clock in today")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.ClockIn(mockToken, "-6.4096", "106.8185")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "you already clock in today")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
func TestClockOut(t *testing.T) {
	data := mocks.NewAttendanceData(t)
	resData := attendance.Core{
		ClockOut:         "07:50",
		ClockOutLocation: "Jalan Soekarno hatta Bandung Jawa Barat",
		ClockOutOSM:      "https://www.openstreetmap.org/#map=16/-6.4096/106.8185",
	}

	t.Run("success ClockOut", func(t *testing.T) {
		data.On("ClockOut", uint(1), "-6.4096", "106.8185").Return(resData, nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.ClockOut(mockToken, "-6.4096", "106.8185")
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.ClockOut, res.ClockOut)
		data.AssertExpectations(t)
	})

	t.Run("server error", func(t *testing.T) {
		data.On("ClockOut", uint(1), "-6.4096", "106.8185").Return(attendance.Core{}, errors.New("server error, location not found")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.ClockOut(mockToken, "-6.4096", "106.8185")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("already ClockOut", func(t *testing.T) {
		data.On("ClockOut", uint(1), "-6.4096", "106.8185").Return(attendance.Core{}, errors.New("user already clock out today")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.ClockOut(mockToken, "-6.4096", "106.8185")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "you already clock out today")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestAttendanceFromAdmin(t *testing.T) {
	data := mocks.NewAttendanceData(t)

	t.Run("success AttendanceFromAdmin", func(t *testing.T) {
		data.On("AttendanceFromAdmin", uint(1), "2023-01-28", "2023-01-28", "annual_leave", uint(1)).Return(nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		err := srv.AttendanceFromAdmin(mockToken, "2023-01-28", "2023-01-28", "annual_leave", uint(1))
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("wrong input", func(t *testing.T) {
		data.On("AttendanceFromAdmin", uint(1), "2023-01-28", "2023-01-28", "annual_leave", uint(1)).Return(errors.New("wrong input format")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		err := srv.AttendanceFromAdmin(mockToken, "2023-01-28", "2023-01-28", "annual_leave", uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong input format")
		data.AssertExpectations(t)
	})
	t.Run("server error", func(t *testing.T) {
		data.On("AttendanceFromAdmin", uint(1), "2023-01-28", "2023-01-28", "annual_leave", uint(1)).Return(errors.New("creating data fail, server error")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		err := srv.AttendanceFromAdmin(mockToken, "2023-01-28", "2023-01-28", "annual_leave", uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		data.AssertExpectations(t)
	})
}

func TestRecord(t *testing.T) {
	data := mocks.NewAttendanceData(t)
	resData := []attendance.Core{
		{
			ClockIn:         "07:50",
			ClockInLocation: "Jalan Soekarno hatta Bandung Jawa Barat",
			ClockInOSM:      "https://www.openstreetmap.org/#map=16/-6.4096/106.8185",
		},
	}
	t.Run("success Record", func(t *testing.T) {
		data.On("Record", uint(1), "2023-01-28", "2023-01-28").Return(resData, nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.Record(mockToken, "2023-01-28", "2023-01-28")
		assert.Equal(t, res, resData)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("wrong input", func(t *testing.T) {
		data.On("Record", uint(1), "2023-01-28", "2023-01-28").Return([]attendance.Core{}, errors.New("wrong input format")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.Record(mockToken, "2023-01-28", "2023-01-28")
		assert.NotNil(t, err)
		assert.Equal(t, []attendance.Core{}, res)
		assert.ErrorContains(t, err, "wrong input format")
		data.AssertExpectations(t)
	})
	t.Run("server error", func(t *testing.T) {
		data.On("Record", uint(1), "2023-01-28", "2023-01-28").Return([]attendance.Core{}, errors.New("data not found")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		mockToken := token.(*jwt.Token)
		mockToken.Valid = true
		res, err := srv.Record(mockToken, "2023-01-28", "2023-01-28")
		assert.NotNil(t, err)
		assert.Equal(t, []attendance.Core{}, res)
		assert.ErrorContains(t, err, "server error")
		data.AssertExpectations(t)
	})
}

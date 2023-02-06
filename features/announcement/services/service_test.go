package services

import (
	"errors"
	"testing"
	"timesync-be/features/announcement"
	"timesync-be/helper"
	"timesync-be/mocks"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestPostAnnouncement(t *testing.T) {
	repo := mocks.NewAnnouncementData(t)
	inputData := announcement.Core{
		ID:      0,
		Nip:     "23001",
		Type:    "personal",
		Title:   "Libur guys",
		Message: "Besok Libur karena ada rapat",
	}
	resData := announcement.Core{ID: uint(1), Nip: "23001", Type: "personal", Title: "Libur guys", Message: "Besok Libur karena ada rapat"}

	t.Run("success post announcement", func(t *testing.T) {
		repo.On("PostAnnouncement", uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.PostAnnouncement(pToken, inputData)
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

		res, err := srv.PostAnnouncement(pToken, inputData)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("PostAnnouncement", uint(1), inputData).Return(announcement.Core{}, errors.New("server problem"))
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.PostAnnouncement(pToken, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestGetAnnouncement(t *testing.T) {
	repo := mocks.NewAnnouncementData(t)
	resData := []announcement.Core{{ID: uint(1), Nip: "23001", Type: "personal", Title: "Libur guys", Message: "Besok Libur karena ada rapat"}}

	t.Run("success get all announcement", func(t *testing.T) {
		repo.On("GetAnnouncement").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetAnnouncement()
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetAnnouncement").Return([]announcement.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		res, err := srv.GetAnnouncement()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetAnnouncement").Return([]announcement.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		res, err := srv.GetAnnouncement()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}

func TestGetAnnouncementDetail(t *testing.T) {
	repo := mocks.NewAnnouncementData(t)
	resData := announcement.Core{
		ID:      uint(1),
		Nip:     "23001",
		Type:    "personal",
		Title:   "Libur guys",
		Message: "Besok Libur karena ada rapat",
	}

	t.Run("success get announcement detail", func(t *testing.T) {
		repo.On("GetAnnouncementDetail", uint(1), uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetAnnouncementDetail(pToken, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.Type, res.Type)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetAnnouncementDetail", uint(1), uint(1)).Return(announcement.Core{}, errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetAnnouncementDetail(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, res.ID, uint(0))
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt token", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetAnnouncementDetail(pToken, uint(1))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetAnnouncementDetail", uint(1), uint(1)).Return(announcement.Core{}, errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetAnnouncementDetail(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.ID, uint(0))
		repo.AssertExpectations(t)
	})

}

func TestDeleteAnnouncement(t *testing.T) {
	repo := mocks.NewAnnouncementData(t)

	t.Run("success delete announcement", func(t *testing.T) {
		repo.On("DeleteAnnouncement", uint(1), uint(1)).Return(nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteAnnouncement(pToken, uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt token", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.DeleteAnnouncement(pToken, uint(1))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("DeleteAnnouncement", uint(1), uint(1)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteAnnouncement(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("DeleteAnnouncement", uint(1), uint(1)).Return(errors.New("server problem")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteAnnouncement(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

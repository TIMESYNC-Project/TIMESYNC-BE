package services

import (
	"errors"
	"log"
	"strings"
	"timesync-be/features/announcement"
	"timesync-be/helper"
)

type announcementUseCase struct {
	qry announcement.AnnouncementData
}

func New(ad announcement.AnnouncementData) announcement.AnnouncementService {
	return &announcementUseCase{
		qry: ad,
	}
}

func (auc *announcementUseCase) PostAnnouncement(token interface{}, newAnnouncement announcement.Core) (announcement.Core, error) {
	userID := helper.ExtractToken(token)

	res, err := auc.qry.PostAnnouncement(uint(userID), newAnnouncement)
	if err != nil {
		log.Println("query error", err.Error())
		return announcement.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

func (auc *announcementUseCase) GetAnnouncement() ([]announcement.Core, error) {
	res, err := auc.qry.GetAnnouncement()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "announcement not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []announcement.Core{}, errors.New(msg)
	}
	return res, nil
}

func (auc *announcementUseCase) GetAnnouncementDetail(token interface{}, announcementID uint) (announcement.Core, error) {
	id := helper.ExtractToken(token)
	res, err := auc.qry.GetAnnouncementDetail(uint(id), announcementID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "user") {
			msg = "user not found"
		} else {
			msg = "announcement not found"
		}
		return announcement.Core{}, errors.New(msg)
	}

	return res, nil
}

func (auc *announcementUseCase) EmployeeInbox(token interface{}) ([]announcement.Core, error) {
	id := helper.ExtractToken(token)
	res, err := auc.qry.EmployeeInbox(uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "announcement not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []announcement.Core{}, errors.New(msg)
	}

	return res, nil

}

func (auc *announcementUseCase) DeleteAnnouncement(token interface{}, announcementID uint) error {
	id := helper.ExtractToken(token)

	err := auc.qry.DeleteAnnouncement(uint(id), announcementID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("data not found")
	}

	return nil
}

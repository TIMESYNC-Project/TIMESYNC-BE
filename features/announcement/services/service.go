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

	if userID <= 0 {
		return announcement.Core{}, errors.New("user not found")
	}

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
			msg = "There is a problem with the server"
		}
		return []announcement.Core{}, errors.New(msg)
	}
	return res, nil
}

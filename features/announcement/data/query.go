package data

import (
	"timesync-be/features/announcement"

	"gorm.io/gorm"
)

type announcementQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) announcement.AnnouncementData {
	return &announcementQuery{
		db: db,
	}
}

func (aq *announcementQuery) PostAnnouncement(adminID uint, newAnnouncement announcement.Core) (announcement.Core, error) {
	cnv := CoreToData(newAnnouncement)
	cnv.UserID = uint(adminID)
	err := aq.db.Create(&cnv).Error
	if err != nil {
		return announcement.Core{}, err
	}

	newAnnouncement.ID = cnv.ID

	return newAnnouncement, nil
}

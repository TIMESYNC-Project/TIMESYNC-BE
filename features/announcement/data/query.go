package data

import (
	"log"
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

func (aq *announcementQuery) GetAnnouncement() ([]announcement.Core, error) {
	res := []Announcement{}
	if err := aq.db.Table("announecements").Joins("JOIN users ON users.id = announcemenets.user_id").Select("announcements.id, announcements.title, announcements.message, announcements.created_at").Find(&res).Error; err != nil {
		log.Println("get all announcement query error : ", err.Error())
		return []announcement.Core{}, err
	}
	result := []announcement.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}
	return result, nil
}

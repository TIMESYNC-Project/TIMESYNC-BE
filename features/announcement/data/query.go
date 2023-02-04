package data

import (
	"errors"
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
	if err := aq.db.Table("announcements").Joins("JOIN users ON users.id = announcements.user_id").Select("announcements.id, announcements.title, announcements.message, announcements.created_at").Find(&res).Error; err != nil {
		log.Println("get all announcement query error : ", err.Error())
		return []announcement.Core{}, err
	}
	result := []announcement.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}
	return result, nil
}

func (aq *announcementQuery) GetAnnouncementDetail(adminID uint, announcementID uint) ([]announcement.Core, error) {
	res := []Announcement{}
	if err := aq.db.Where("id = ?", announcementID).Find(&res).Error; err != nil {
		log.Println("get announcement by id query error : ", err.Error())
		return []announcement.Core{}, err
	}
	result := []announcement.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}
	return result, nil
}

func (aq *announcementQuery) DeleteAnnouncement(adminID uint, announcementID uint) error {
	getID := Announcement{}
	err := aq.db.Where("id = ?", announcementID).First(&getID).Error
	if err != nil {
		log.Println("get announcement error : ", err.Error())
		return errors.New("failed to get announcement data")
	}

	qryDelete := aq.db.Delete(&Announcement{}, announcementID)
	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete announcement, data not found")
	}

	return nil
}

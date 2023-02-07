package data

import (
	"errors"
	"fmt"
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
	data := CoreToData(newAnnouncement)
	if newAnnouncement.Nip != "" {
		//query untuk mencari userID berdasarkan NIP
		userData := User{}
		err := aq.db.Where("nip = ?", newAnnouncement.Nip).First(&userData).Error
		if err != nil {
			log.Println("find nip query error")
			return announcement.Core{}, errors.New("employee not found")
		}
		data.UserID = userData.ID
	}

	//logic untuk menentukan type pada saat post announcement
	if newAnnouncement.Nip == "" {
		data.Type = "public"
	} else {
		data.Type = "personal"
	}

	err := aq.db.Create(&data).Error
	if err != nil {
		return announcement.Core{}, err
	}

	newAnnouncement.ID = data.ID
	newAnnouncement.Type = data.Type

	return newAnnouncement, nil
}

func (aq *announcementQuery) GetAnnouncement() ([]announcement.Core, error) {
	res := []Announcement{}
	if err := aq.db.Find(&res).Error; err != nil {
		log.Println("get all announcement query error : ", err.Error())
		return []announcement.Core{}, err
	}
	result := []announcement.Core{}
	i := 0
	for _, val := range res {
		result = append(result, ToCore(val))
		if res[i].Type == "personal" {
			user := User{}
			if err := aq.db.Where("id = ?", res[i].UserID).First(&user).Error; err != nil {
				log.Println("get user by id query error : ", err.Error())
				return []announcement.Core{}, err
			}
			result[i].Name = user.Name
			result[i].Nip = user.Nip
			y := res[i].CreatedAt.Year()
			m := int(res[i].CreatedAt.Month())
			d := res[i].CreatedAt.Day()
			result[i].AnnouncementDate = fmt.Sprintf("%d-%d-%d", y, m, d)
		} else {
			y := res[i].CreatedAt.Year()
			m := int(res[i].CreatedAt.Month())
			d := res[i].CreatedAt.Day()
			result[i].AnnouncementDate = fmt.Sprintf("%d-%d-%d", y, m, d)
		}

		i++
	}

	return result, nil
}

func (aq *announcementQuery) GetAnnouncementDetail(adminID uint, announcementID uint) (announcement.Core, error) {
	res := Announcement{}
	if err := aq.db.Where("id = ?", announcementID).First(&res).Error; err != nil {
		log.Println("get announcement by id query error : ", err.Error())
		return announcement.Core{}, err
	}

	user := User{}
	if err := aq.db.Where("id = ?", res.UserID).First(&user).Error; err != nil {
		log.Println("get user by id query error : ", err.Error())
		return announcement.Core{}, err
	}

	y := res.CreatedAt.Year()
	m := int(res.CreatedAt.Month())
	d := res.CreatedAt.Day()
	result := ToCore(res)
	result.AnnouncementDate = fmt.Sprintf("%d-%d-%d", y, m, d)
	result.Name = user.Name
	result.Nip = user.Nip
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

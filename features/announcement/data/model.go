package data

import (
	"timesync-be/features/announcement"

	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model
	UserID  uint
	Type    string
	Title   string
	Message string
}

type User struct {
	gorm.Model
	ProfilePicture string
	Name           string
	Nip            string
	Position       string
	Role           string
}

func ToCore(data Announcement) announcement.Core {
	return announcement.Core{
		ID:      data.ID,
		Type:    data.Type,
		Title:   data.Title,
		Message: data.Message,
	}
}

func CoreToData(data announcement.Core) Announcement {
	return Announcement{
		Model:   gorm.Model{ID: data.ID},
		Type:    data.Type,
		Title:   data.Title,
		Message: data.Message,
	}
}

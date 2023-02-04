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

func ToCore(data Announcement) announcement.Core {
	return announcement.Core{
		ID:      data.ID,
		UserID:  data.UserID,
		Type:    data.Type,
		Title:   data.Title,
		Message: data.Message,
	}
}

func CoreToData(data announcement.Core) Announcement {
	return Announcement{
		Model:   gorm.Model{ID: data.ID},
		UserID:  data.UserID,
		Type:    data.Type,
		Title:   data.Title,
		Message: data.Message,
	}
}

func (dataModel *Announcement) AllModelsToCore() announcement.Core {
	return announcement.Core{
		ID:      dataModel.ID,
		UserID:  dataModel.UserID,
		Type:    dataModel.Type,
		Title:   dataModel.Title,
		Message: dataModel.Message,
	}
}

func AllListToCore(data []Announcement) []announcement.Core {
	var dataCore []announcement.Core
	for _, v := range data {
		dataCore = append(dataCore, v.AllModelsToCore())
	}
	return dataCore
}

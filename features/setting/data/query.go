package data

import (
	"errors"
	"log"
	"timesync-be/features/setting"

	"gorm.io/gorm"
)

type settingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) setting.SettingData {
	return &settingQuery{
		db: db,
	}
}

// GetSetting implements setting.SettingData
func (sq *settingQuery) GetSetting() (setting.Core, error) {
	data := Setting{}
	err := sq.db.First(&data).Error
	if err != nil {
		log.Println("query error")
		return setting.Core{}, errors.New("setting not found")
	}
	return DataToCore(data), nil
}

// EditSetting implements setting.SettingData
func (sq *settingQuery) EditSetting(userID uint, updateSetting setting.Core) (setting.Core, error) {
	upd := CoreToData(updateSetting)
	err := sq.db.Where("id = ?", 1).Updates(&upd).Error
	if err != nil {
		log.Println("query error")
		return setting.Core{}, errors.New("update fail")
	}
	return updateSetting, nil
}

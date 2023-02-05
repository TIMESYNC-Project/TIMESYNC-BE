package service

import (
	"errors"
	"log"
	"timesync-be/features/setting"
	"timesync-be/helper"
)

type settingUseCase struct {
	qry setting.SettingData
}

func New(sd setting.SettingData) setting.SettingService {
	return &settingUseCase{
		qry: sd,
	}
}

// GetSetting implements setting.SettingService
func (suc *settingUseCase) GetSetting() (setting.Core, error) {
	res, err := suc.qry.GetSetting()
	if err != nil {
		log.Println("data not found")
		return setting.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// EditSetting implements setting.SettingService
func (suc *settingUseCase) EditSetting(token interface{}, updateSetting setting.Core) (setting.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := suc.qry.EditSetting(uint(userID), updateSetting)
	if err != nil {
		log.Println("data not found")
		return setting.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

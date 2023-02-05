package setting

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint   `json:"id"`
	Start       string `json:"working_hour_start"`
	End         string `json:"working_hour_end"`
	Tolerance   int    `json:"tolerance"`
	AnnualLeave int    `json:"annual_leave"`
}

type SettingHandler interface {
	GetSetting() echo.HandlerFunc
	EditSetting() echo.HandlerFunc
}

type SettingService interface {
	GetSetting() (Core, error)
	EditSetting(token interface{}, updateSetting Core) (Core, error)
}

type SettingData interface {
	GetSetting() (Core, error)
	EditSetting(userID uint, updateSetting Core) (Core, error)
}

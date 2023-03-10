package handler

import (
	"net/http"
	"strings"
	"timesync-be/features/setting"

	"github.com/labstack/echo/v4"
)

type settingController struct {
	srv setting.SettingService
}

func New(srv setting.SettingService) setting.SettingHandler {
	return &settingController{
		srv: srv,
	}
}

// GetSetting implements setting.SettingHandler
func (sc *settingController) GetSetting() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := sc.srv.GetSetting()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show setting",
		})
	}
}

// EditSetting implements setting.SettingHandler
func (sc *settingController) EditSetting() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := EditSetting{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		res, err := sc.srv.EditSetting(c.Get("user"), *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "access denied") {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success change setting",
		})
	}
}

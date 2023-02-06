package handler

import (
	"net/http"
	"timesync-be/features/company"

	"github.com/labstack/echo/v4"
)

type companyController struct {
	srv company.CompanyService
}

func New(cs company.CompanyService) company.CompanyHandler {
	return &companyController{
		srv: cs,
	}
}

// EditProfile implements company.CompanyHandler
func (cc *companyController) EditProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := EditRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("file")
		if checkFile != nil {
			formHeader, err := c.FormFile("file")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		res, err := cc.srv.EditProfile(c.Get("user"), input.FileHeader, *ReqToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show company profile",
		})
	}
}

// GetProfile implements company.CompanyHandler
func (cc *companyController) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cc.srv.GetProfile()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show company profile",
		})
	}
}
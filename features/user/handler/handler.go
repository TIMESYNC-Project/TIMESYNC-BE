package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"timesync-be/features/user"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		res, err := uc.srv.Register(*ReqToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create account"})
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		if input.Nip == "" {
			return c.JSON(http.StatusBadRequest, "nip not allowed empty")
		} else if input.Password == "" {
			return c.JSON(http.StatusBadRequest, "password not allowed empty")
		}

		token, res, err := uc.srv.Login(input.Nip, input.Password)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "success login", ToResponse(res), token))
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uc.srv.Delete(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error, account fail to delete",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete profile",
		})
	}
}

// Update implements user.UserHandler
func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		eID := c.Param("id")
		employeeID, _ := strconv.Atoi(eID)
		input := RegisterRequest{}
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
		res, err := uc.srv.Update(uint(employeeID), input.FileHeader, *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already used"})
			} else if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusOK, map[string]interface{}{
			// "data":    res,
			"message": "update profile success",
		})
	}
}

// Csv implements user.UserHandler
func (uc *userControll) Csv() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Select a file to upload"})
		}
		res, err := uc.srv.Csv(*formHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create account from csv"})
	}
}

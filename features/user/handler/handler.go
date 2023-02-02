package handler

import (
	"log"
	"net/http"
	"timesync-be/features/user"

	"github.com/labstack/echo"
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

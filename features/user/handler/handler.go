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
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		res, err := uc.srv.Register(c.Get("user"), *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "access") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
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
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "nip not allowed empty"})
		} else if input.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not allowed empty"})
		}

		expired, token, res, err := uc.srv.Login(input.Nip, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not match"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":             ToResponse(res),
			"message":          "success login",
			"token":            token,
			"token_expired_at": expired,
		})
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		employeeID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}
		err = uc.srv.Delete(token, uint(employeeID))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "data not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success deactivate employee profile",
		})
	}
}

// Update implements user.UserHandler
func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("profile_picture")
		if checkFile != nil {
			formHeader, err := c.FormFile("profile_picture")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		// log.Println((input.FileHeader))
		res, err := uc.srv.Update(c.Get("user"), input.FileHeader, *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already used"})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "access denied") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "access denied"})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "file size max 500kb"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "not registered") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})
			}
		}

		result, err := ConvertEmployeeUpdateResponse(res)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			// log.Println(res)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    result,
				"message": "success update employee profile",
			})
		}

	}
}

// Csv implements user.UserHandler
func (uc *userControll) Csv() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Select a file to upload"})
		}

		err = uc.srv.Csv(*formHeader)
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create account from csv"})
	}
}

// Profile implements user.UserHandler
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// eID := c.Param("id")
		// employeeID, _ := strconv.Atoi(eID)
		res, err := uc.srv.Profile(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToProfileResponse(res),
			"message": "success show profile",
		})
	}
}

// ProfileEmployee implements user.UserHandler
func (uc *userControll) ProfileEmployee() echo.HandlerFunc {
	return func(c echo.Context) error {
		eID := c.Param("id")
		employeeID, _ := strconv.Atoi(eID)
		res, err := uc.srv.ProfileEmployee(uint(employeeID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToProfileResponse(res),
			"message": "success show profile",
		})
	}
}

// AdminEditEmployee implements user.UserHandler
func (uc *userControll) AdminEditEmployee() echo.HandlerFunc {
	return func(c echo.Context) error {
		eID := c.Param("id")
		employeeID, _ := strconv.Atoi(eID)
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("profile_picture")
		if checkFile != nil {
			formHeader, err := c.FormFile("profile_picture")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		res, err := uc.srv.AdminEditEmployee(c.Get("user"), uint(employeeID), input.FileHeader, *ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email duplicated") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already used"})
			} else if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "access") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		result, err := ConvertUpdateResponse(res)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			// log.Println(res)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    result,
				"message": "update profile success",
			})
		}
	}
}

func (uc *userControll) GetAllEmployee() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.srv.GetAllEmployee()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		result := []ShowAllEmployee{}
		for _, val := range res {
			result = append(result, ShowAllEmployeeJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all employee",
		})
	}
}

// Search implements user.UserHandler
func (uc *userControll) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		quotes := c.QueryParam("q")
		if quotes == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "nothing to search"})
		}
		res, err := uc.srv.Search(c.Get("user"), quotes)
		if err != nil {
			if strings.Contains(err.Error(), "access denied") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		}
		if len(res) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "user not found"})
		}
		result := []Search{}
		for i := 0; i < len(res); i++ {
			result = append(result, SearchResponse(res[i]))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "searching success",
		})
	}
}

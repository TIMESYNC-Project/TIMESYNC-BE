package handler

import (
	"net/http"
	"strings"
	"timesync-be/features/user"
)

type RegResp struct {
}

func ToRegResp(data user.Core) RegResp {
	return RegResp{}
}

type UserReponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nip      string `json:"nip"`
	Gender   string `json:"gender"`
	Position string `json:"position"`
	Role     string `json:"role"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Nip:      data.Nip,
		Gender:   data.Nip,
		Position: data.Position,
		Role:     data.Role,
	}
}

type UpdateUserResp struct {
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Nip            string `json:"nip"`
	Email          string `json:"email"`
	Position       string `json:"position"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Password       string `json:"password"`
}

func ToResponseUpd(data user.Core) UpdateUserResp {
	return UpdateUserResp{
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		Nip:            data.Nip,
		Email:          data.Email,
		Position:       data.Position,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = (data[0])
	} else {
		resp["data"] = (data[0])
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	}

	return code, resp
}

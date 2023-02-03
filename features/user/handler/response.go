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
		Gender:   data.Gender,
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

type ProfileResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	BirthOfDate    string `json:"birth_of_date"`
	Nip            string `json:"nip"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Position       string `json:"position"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	AnnualLeave    int    `json:"annual_leave"`
}

func ToProfileResponse(data user.Core) ProfileResponse {
	return ProfileResponse{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		BirthOfDate:    data.BirthOfDate,
		Email:          data.Email,
		Nip:            data.Nip,
		Gender:         data.Gender,
		Position:       data.Position,
		Phone:          data.Phone,
		Address:        data.Address,
		AnnualLeave:    data.AnnualLeave,
	}
}

type UpdateResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name" form:"name"`
	BirthOfDate    string `json:"birth_of_date" form:"birth_of_date"`
	Email          string `json:"email" form:"email"`
	Gender         string `json:"gender" form:"gender"`
	Position       string `json:"position" form:"position"`
	Phone          string `json:"phone" form:"phone"`
	Address        string `json:"address" form:"address"`
	Password       string `json:"password" form:"password"`
}

func ToUpdateResponse(data user.Core) UpdateResponse {
	return UpdateResponse{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		BirthOfDate:    data.BirthOfDate,
		Email:          data.Email,
		Gender:         data.Gender,
		Position:       data.Position,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
	}
}

type UpdateResponseEmployee struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Password       string `json:"password" form:"password"`
}

func ToUpdateResponseEmployee(data user.Core) UpdateResponseEmployee {
	return UpdateResponseEmployee{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Password:       data.Password,
	}
}

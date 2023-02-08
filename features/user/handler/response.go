package handler

import (
	"errors"
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

type ShowAllEmployee struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Nip            string `json:"nip"`
	Position       string `json:"position"`
}

func ShowAllEmployeeJson(data user.Core) ShowAllEmployee {
	return ShowAllEmployee{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		Nip:            data.Nip,
		Position:       data.Position,
	}
}

type Search struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Nip            string `json:"nip"`
	Position       string `json:"position"`
}

func SearchResponse(data user.Core) Search {
	return Search{
		ID:       data.ID,
		Name:     data.Name,
		Nip:      data.Nip,
		Position: data.Position,
	}
}
func ConvertUpdateResponse(inputan user.Core) (interface{}, error) {
	ResponseFilter := user.Core{}
	ResponseFilter = inputan
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.ProfilePicture != "" {
		result["profile_picture"] = ResponseFilter.ProfilePicture
	}
	if ResponseFilter.Name != "" {
		result["name"] = ResponseFilter.Name
	}
	if ResponseFilter.BirthOfDate != "" {
		result["birth_of_date"] = ResponseFilter.BirthOfDate
	}
	if ResponseFilter.Email != "" {
		result["email"] = ResponseFilter.Email
	}
	if ResponseFilter.Gender != "" {
		result["gender"] = ResponseFilter.Gender
	}
	if ResponseFilter.Position != "" {
		result["position"] = ResponseFilter.Position
	}
	if ResponseFilter.Phone != "" {
		result["phone"] = ResponseFilter.Phone
	}
	if ResponseFilter.Address != "" {
		result["address"] = ResponseFilter.Address
	}
	if ResponseFilter.Password != "" {
		result["password"] = ResponseFilter.Password
	}

	if len(result) <= 1 {
		return user.Core{}, errors.New("no data was change")
	}
	return result, nil
}

// ID:             data.ID,
// ProfilePicture: data.ProfilePicture,
// Name:           data.Name,
// BirthOfDate:    data.BirthOfDate,
// Email:          data.Email,
// Gender:         data.Gender,
// Position:       data.Position,
// Phone:          data.Phone,
// Address:        data.Address,
// Password:       data.Password,

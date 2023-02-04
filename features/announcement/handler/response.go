package handler

import (
	"net/http"
	"strings"
	"time"
	"timesync-be/features/announcement"
)

type PostAnnouncementReponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func ToPostAnnouncementReponse(data announcement.Core) PostAnnouncementReponse {
	return PostAnnouncementReponse{
		ID:        data.ID,
		Title:     data.Title,
		Message:   data.Message,
		CreatedAt: time.Now(),
	}
}

func CoresToResponse(dataCore announcement.Core) PostAnnouncementReponse {
	return PostAnnouncementReponse{
		ID:      dataCore.ID,
		Title:   dataCore.Title,
		Message: dataCore.Message,
	}
}

func ListCoreToResp(data []announcement.Core) []PostAnnouncementReponse {
	var dataResp []PostAnnouncementReponse
	for _, v := range data {
		dataResp = append(dataResp, CoresToResponse(v))
	}
	return dataResp
}

type ShowAllAnnouncement struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func ShowAllAnnouncementJson(data announcement.Core) ShowAllAnnouncement {
	return ShowAllAnnouncement{
		ID:      data.ID,
		Title:   data.Title,
		Message: data.Message,
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

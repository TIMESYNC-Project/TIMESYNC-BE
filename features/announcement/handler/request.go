package handler

import "timesync-be/features/announcement"

type PostAnnouncementRequest struct {
	UserID  uint   `json:"user_id" form:"user_id"`
	Type    string `json:"type" form:"type"`
	Title   string `json:"title" form:"title"`
	Message string `json:"message" form:"message"`
}

func ReqToCore(data interface{}) *announcement.Core {
	res := announcement.Core{}

	switch data.(type) {
	case PostAnnouncementRequest:
		cnv := data.(PostAnnouncementRequest)
		res.UserID = cnv.UserID
		res.Type = cnv.Type
		res.Title = cnv.Title
		res.Message = cnv.Message
	default:
		return nil
	}

	return &res
}

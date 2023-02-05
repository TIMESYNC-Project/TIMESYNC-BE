package handler

import "timesync-be/features/announcement"

type PostAnnouncementRequest struct {
	Nip     string `json:"to" form:"to"`
	Type    string `json:"type" form:"type"`
	Title   string `json:"announcement_title" form:"announcement_title"`
	Message string `json:"announcement_description" form:"announcement_description"`
}

func ReqToCore(data interface{}) *announcement.Core {
	res := announcement.Core{}

	switch data.(type) {
	case PostAnnouncementRequest:
		cnv := data.(PostAnnouncementRequest)
		res.Nip = cnv.Nip
		res.Title = cnv.Title
		res.Message = cnv.Message
	default:
		return nil
	}

	return &res
}

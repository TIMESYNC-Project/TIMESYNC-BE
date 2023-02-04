package handler

import "timesync-be/features/announcement"

type PostAnnouncementRequest struct {
	Nip     string `json:"nip" form:"nip"`
	Type    string `json:"type" form:"type"`
	Title   string `json:"title" form:"title"`
	Message string `json:"message" form:"message"`
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

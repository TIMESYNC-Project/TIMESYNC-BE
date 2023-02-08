package handler

import (
	"mime/multipart"
	"timesync-be/features/approval"
)

type PostApprovalRequest struct {
	Title       string `json:"approval_title" form:"approval_title"`
	StartDate   string `json:"approval_start_date" form:"approval_start_date"`
	EndDate     string `json:"approval_end_date" form:"approval_end_date"`
	Description string `json:"approval_description" form:"approval_description"`
	FileHeader  multipart.FileHeader
}

type UpdateApprovalRequest struct {
	Status string `json:"approval_status" form:"approval_status"`
}

func ReqToCore(data interface{}) *approval.Core {
	res := approval.Core{}

	switch data.(type) {
	case PostApprovalRequest:
		cnv := data.(PostApprovalRequest)
		res.Title = cnv.Title
		res.StartDate = cnv.StartDate
		res.EndDate = cnv.EndDate
		res.Description = cnv.Description
	case UpdateApprovalRequest:
		cnv := data.(UpdateApprovalRequest)
		res.Status = cnv.Status
	default:
		return nil
	}

	return &res
}

package handler

import (
	"mime/multipart"
	"timesync-be/features/approval"
)

type PostApprovalRequest struct {
	Title         string `json:"approval_title" form:"approval_title"`
	StartDate     string `json:"approval_start_date" form:"approval_start_date"`
	EndDate       string `json:"approval_end_date" form:"approval_end_date"`
	Description   string `json:"approval_description" form:"approval_description"`
	ApprovalImage string `json:"approval_image" form:"approval_image"`
	FileHeader    multipart.FileHeader
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
		res.ApprovalImage = cnv.ApprovalImage
	default:
		return nil
	}

	return &res
}

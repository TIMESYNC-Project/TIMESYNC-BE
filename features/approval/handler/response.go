package handler

import "timesync-be/features/approval"

type PostApprovalResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"approval_title"`
	StartDate     string `json:"approval_start_date"`
	EndDate       string `json:"approval_end_date"`
	Description   string `json:"approval_description"`
	ApprovalImage string `json:"approval_image"`
	Status        string `json:"approval_status"`
}

type ShowAllApproval struct {
	ID      uint   `json:"id"`
	Title   string `json:"approval_title"`
	EndDate string `json:"approval_end_date"`
	Status  string `json:"approval_status"`
}

func ToPostApprovalResponse(data approval.Core) PostApprovalResponse {
	return PostApprovalResponse{
		ID:            data.ID,
		Title:         data.Title,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		Description:   data.Description,
		ApprovalImage: data.ApprovalImage,
		Status:        data.Status,
	}
}

func ShowAllApprovalJson(data approval.Core) ShowAllApproval {
	return ShowAllApproval{
		ID:      data.ID,
		Title:   data.Title,
		EndDate: data.EndDate,
		Status:  data.Status,
	}
}

package data

import (
	"timesync-be/features/approval"

	"gorm.io/gorm"
)

type Approval struct {
	gorm.Model
	UserID        uint
	Title         string
	StartDate     string
	EndDate       string
	Description   string
	ApprovalImage string
	Status        string
}

type User struct {
	gorm.Model
	ProfilePicture string
	Name           string
	Nip            string
	Position       string
	Role           string
}

func ToCore(data Approval) approval.Core {
	return approval.Core{
		ID:            data.ID,
		Title:         data.Title,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		Description:   data.Description,
		ApprovalImage: data.ApprovalImage,
		Status:        data.Status,
	}
}

func CoreToData(data approval.Core) Approval {
	return Approval{
		Model:         gorm.Model{ID: data.ID},
		Title:         data.Title,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		Description:   data.Description,
		ApprovalImage: data.ApprovalImage,
		Status:        data.Status,
	}
}

package data

import (
	"errors"
	"timesync-be/features/approval"

	"gorm.io/gorm"
)

type approvalQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) approval.ApprovalData {
	return &approvalQuery{
		db: db,
	}
}

func (aq *approvalQuery) PostApproval(employeeID uint, newApproval approval.Core) (approval.Core, error) {
	cnv := CoreToData(newApproval)
	cnv.UserID = uint(employeeID)
	cnv.Status = "pending"
	err := aq.db.Create(&cnv).Error
	if err != nil {
		return approval.Core{}, errors.New("create new approval query error")
	}

	newApproval.ID = cnv.ID
	newApproval.Status = cnv.Status

	return newApproval, nil
}

// GetApproval implements approval.ApprovalData
func (*approvalQuery) GetApproval() ([]approval.Core, error) {
	panic("unimplemented")
}

// UpdateApproval implements approval.ApprovalData
func (*approvalQuery) UpdateApproval(adminID uint, approvalID uint) ([]approval.Core, error) {
	panic("unimplemented")
}

package data

import (
	"errors"
	"fmt"
	"log"
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
func (aq *approvalQuery) GetApproval() ([]approval.Core, error) {
	res := []Approval{}
	if err := aq.db.Table("approvals").Joins("JOIN users ON users.id = approvals.user_id").Select("approvals.id, users.name,approvals.title, approvals.start_date, approvals.end_date, approvals.status").Find(&res).Error; err != nil {
		log.Println("get all approvals record query error : ", err.Error())
		return []approval.Core{}, errors.New("get all approval query error")
	}
	result := []approval.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}
	return result, nil
}

func (aq *approvalQuery) ApprovalDetail(approvalID uint) (approval.Core, error) {
	res := Approval{}
	if err := aq.db.Where("id = ?", approvalID).First(&res).Error; err != nil {
		log.Println("get detail approval query error : ", err.Error())
		return approval.Core{}, errors.New("get detail approval query error")
	}
	result := ToCore(res)
	user := User{}
	if res.UserID != 0 {
		if err := aq.db.Where("id = ?", res.UserID).First(&user).Error; err != nil {
			log.Println("get user by id query error : ", err.Error())
			return approval.Core{}, errors.New("get user by id error")
		}
		result.Name = user.Name
	}
	y := res.CreatedAt.Year()
	m := int(res.CreatedAt.Month())
	d := res.CreatedAt.Day()
	result.ApprovalDate = fmt.Sprintf("%d-%d-%d", y, m, d)

	return result, nil
}

// UpdateApproval implements approval.ApprovalData
func (aq *approvalQuery) UpdateApproval(adminID uint, approvalID uint, updatedApproval approval.Core) (approval.Core, error) {
	getID := Approval{}
	err := aq.db.Where("id = ?", approvalID).First(&getID).Error
	if err != nil {
		log.Println("get approval error : ", err.Error())
		return approval.Core{}, errors.New("get approval query error")
	}

	if adminID != 1 {
		log.Println("Unauthorized request")
		return approval.Core{}, errors.New("unauthorized request")
	}

	cnv := CoreToData(updatedApproval)
	qry := aq.db.Where("id = ?", approvalID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update approval query error : data not found")
		return approval.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update approval query error : ", err.Error())
	}

	updatedApproval.ID = getID.ID
	updatedApproval.Title = getID.Title
	updatedApproval.StartDate = getID.StartDate
	updatedApproval.EndDate = getID.EndDate
	updatedApproval.Description = getID.Description
	updatedApproval.ApprovalImage = getID.ApprovalImage

	return updatedApproval, nil
}

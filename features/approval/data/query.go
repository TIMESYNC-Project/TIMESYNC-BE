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
		log.Println("query error", err.Error())
		return approval.Core{}, errors.New("server error, failed to query")
	}

	newApproval.ID = cnv.ID
	newApproval.Status = cnv.Status

	return newApproval, nil
}

// GetApproval implements approval.ApprovalData
func (aq *approvalQuery) GetApproval() ([]approval.Core, error) {
	res := []Approval{}
	if err := aq.db.Order("create_at DESC").Find(&res).Error; err != nil {
		log.Println("get all approvals record query error : ", err.Error())
		return []approval.Core{}, errors.New("get all approval query error")
	}
	i := 0
	result := []approval.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
		y := val.CreatedAt.Year()
		m := int(val.CreatedAt.Month())
		d := val.CreatedAt.Day()
		result[i].ApprovalDate = fmt.Sprintf("%d-%d-%d", y, m, d)
		user := User{}
		if val.UserID != 0 {
			if err := aq.db.Where("id = ?", val.UserID).First(&user).Error; err != nil {
				result[i].Name = "user was deactivated"
			} else {
				result[i].Name = user.Name
			}

		}
		i++
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
	result.Description = res.Description

	return result, nil
}

func (aq *approvalQuery) EmployeeApprovalRecord(employeeID uint) ([]approval.Core, error) {
	res := []Approval{}
	if err := aq.db.Where("user_id = ?", employeeID).Find(&res).Error; err != nil {
		log.Println("get employee approval query error : ", err.Error())
		return []approval.Core{}, err
	}
	i := 0
	result := []approval.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
		y := val.CreatedAt.Year()
		m := int(val.CreatedAt.Month())
		d := val.CreatedAt.Day()
		result[i].ApprovalDate = fmt.Sprintf("%d-%d-%d", y, m, d)
		user := User{}
		if val.UserID != 0 {
			if err := aq.db.Where("id = ?", val.UserID).First(&user).Error; err != nil {
				log.Println("get user by id query error : ", err.Error())
				return []approval.Core{}, errors.New("get user by id error")
			}
			result[i].Name = user.Name
		}
		i++
	}

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

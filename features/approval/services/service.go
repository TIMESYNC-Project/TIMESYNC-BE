package services

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"timesync-be/features/approval"
	"timesync-be/helper"
)

type approvalUseCase struct {
	qry approval.ApprovalData
}

func New(ad approval.ApprovalData) approval.ApprovalService {
	return &approvalUseCase{
		qry: ad,
	}
}

func (auc *approvalUseCase) PostApproval(token interface{}, fileData multipart.FileHeader, newApproval approval.Core) (approval.Core, error) {
	employeeID := helper.ExtractToken(token)

	//validation
	err := helper.ApprovalValidation(newApproval)
	if err != nil {
		return approval.Core{}, errors.New("validate: " + err.Error())
	}

	// kondisi dibawah dilakukan agar foto bisa kosong dan agar unit testing tidak error
	if fileData.Size != 0 {
		res, err := helper.GetUrlImagesFromAWS(fileData)
		if err != nil {
			return approval.Core{}, errors.New("validate: " + err.Error())
		}
		newApproval.ApprovalImage = res
	}
	res, err := auc.qry.PostApproval(uint(employeeID), newApproval)
	if err != nil {
		log.Println("cannot create approval", err.Error())
		return approval.Core{}, errors.New("server error")
	}
	return res, nil
}

// GetApproval implements approval.ApprovalService
func (auc *approvalUseCase) GetApproval() ([]approval.Core, error) {
	res, err := auc.qry.GetApproval()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "approval not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []approval.Core{}, errors.New(msg)
	}
	return res, nil
}

func (auc *approvalUseCase) ApprovalDetail(approvalID uint) (approval.Core, error) {
	res, err := auc.qry.ApprovalDetail(approvalID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "approval not found"
		} else {
			msg = "there is a problem with the server"
		}
		return approval.Core{}, errors.New(msg)
	}
	return res, nil
}

func (auc *approvalUseCase) EmployeeApprovalRecord(token interface{}) ([]approval.Core, error) {
	id := helper.ExtractToken(token)
	res, err := auc.qry.EmployeeApprovalRecord(uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "approval not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []approval.Core{}, errors.New(msg)
	}

	return res, nil
}

// UpdateApproval implements approval.ApprovalService
func (auc *approvalUseCase) UpdateApproval(token interface{}, approvalID uint, updatedApproval approval.Core) (approval.Core, error) {
	adminID := helper.ExtractToken(token)
	res, err := auc.qry.UpdateApproval(uint(adminID), approvalID, updatedApproval)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "failed to update, no new record or data not found"
		} else if strings.Contains(err.Error(), "unauthorized") {
			msg = "unauthorized request"
		} else {
			msg = "unable to process the data"
		}
		return approval.Core{}, errors.New(msg)
	}
	res.ID = approvalID
	res.UserID = uint(adminID)

	return res, nil

}

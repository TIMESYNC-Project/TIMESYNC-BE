package services

import (
	"errors"
	"mime/multipart"
	"strings"
	"timesync-be/features/approval"
	"timesync-be/helper"

	uuid "github.com/satori/go.uuid"
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

	if fileData.Filename != "" {
		if fileData.Size > 500000 {
			return approval.Core{}, errors.New("size error")
		}
		file, err := fileData.Open()
		if err != nil {
			return approval.Core{}, errors.New("error open fileData")
		}
		defer file.Close()
		// Validasi Type
		_, err = helper.TypeFile(file)
		if err != nil {
			return approval.Core{}, errors.New("file type error only jpg or png file can be upload")
		}
		fileName := uuid.NewV4().String()
		fileData.Filename = fileName + fileData.Filename[(len(fileData.Filename)-5):len(fileData.Filename)]
		src, _ := fileData.Open()
		defer src.Close()
		uploadURL, err := helper.UploadToS3(fileData.Filename, src)
		if err != nil {
			return approval.Core{}, errors.New("cannot upload to s3 server error")
		}
		newApproval.ApprovalImage = uploadURL
	}
	res, err := auc.qry.PostApproval(uint(employeeID), newApproval)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server error"
		}
		return approval.Core{}, errors.New(msg)
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

	if id <= 0 {
		return []approval.Core{}, errors.New("data not found")
	}
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

	if adminID <= 0 {
		return approval.Core{}, errors.New("data not found")
	}

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

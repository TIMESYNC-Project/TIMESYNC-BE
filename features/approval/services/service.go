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

	if employeeID <= 0 {
		return approval.Core{}, errors.New("data not found")
	}

	if fileData.Size != 0 {
		if fileData.Size > 500000 {
			return approval.Core{}, errors.New("size error")
		}
		fileName := uuid.NewV4().String()
		fileData.Filename = fileName + fileData.Filename[(len(fileData.Filename)-5):len(fileData.Filename)]
		src, err := fileData.Open()
		if err != nil {
			return approval.Core{}, errors.New("error open fileData")
		}
		// Validasi Type
		if !helper.TypeFile(src) {
			return approval.Core{}, errors.New("file type error")
		}
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
			msg = "Failed to update, no new record or data not found"
		} else if strings.Contains(err.Error(), "Unauthorized") {
			msg = "Unauthorized request"
		} else {
			msg = "unable to process the data"
		}
		return approval.Core{}, errors.New(msg)
	}
	res.ID = approvalID
	res.UserID = uint(adminID)

	return res, nil

}

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
func (*approvalUseCase) UpdateApproval(token interface{}, approvalID uint) ([]approval.Core, error) {
	panic("unimplemented")
}

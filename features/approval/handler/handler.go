package handler

import (
	"log"
	"net/http"
	"timesync-be/features/approval"

	"github.com/labstack/echo/v4"
)

type approvalControll struct {
	srv approval.ApprovalService
}

func New(srv approval.ApprovalService) approval.ApprovalHandler {
	return &approvalControll{
		srv: srv,
	}
}

// PostApproval implements approval.ApprovalHandler
func (ac *approvalControll) PostApproval() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := PostApprovalRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		checkFile, _, _ := c.Request().FormFile("file")
		if checkFile != nil {
			formHeader, err := c.FormFile("file")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		res, err := ac.srv.PostApproval(c.Get("user"), input.FileHeader, *ReqToCore(input))
		if err != nil {
			log.Println("error post approval : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ToPostApprovalResponse(res),
			"message": "send an approval success",
		})
	}
}

// GetApproval implements approval.ApprovalHandler
func (*approvalControll) GetApproval() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateApproval implements approval.ApprovalHandler
func (*approvalControll) UpdateApproval() echo.HandlerFunc {
	panic("unimplemented")
}

package handler

import (
	"log"
	"net/http"
	"strconv"
	"timesync-be/features/approval"
	"timesync-be/helper"

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
func (ac *approvalControll) GetApproval() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.srv.GetApproval()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []ShowAllApproval{}
		for _, val := range res {
			result = append(result, ShowAllApprovalJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show employee approval record",
		})

	}
}

// UpdateApproval implements approval.ApprovalHandler
func (ac *approvalControll) UpdateApproval() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")

		approvalID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		input := UpdateApprovalRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := ac.srv.UpdateApproval(token, uint(approvalID), *ReqToCore(input))

		if err != nil {
			log.Println("error update approval : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ToPostApprovalResponse(res),
			"message": "success approve employee permission",
		})
	}

}

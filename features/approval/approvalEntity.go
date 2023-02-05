package approval

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	UserID        uint
	Title         string
	StartDate     string
	EndDate       string
	Description   string
	ApprovalImage string
	Status        string
}

type ApprovalHandler interface {
	PostApproval() echo.HandlerFunc
	GetApproval() echo.HandlerFunc
	UpdateApproval() echo.HandlerFunc
}

type ApprovalService interface {
	PostApproval(token interface{}, fileData multipart.FileHeader, newApproval Core) (Core, error)
	GetApproval() ([]Core, error)
	UpdateApproval(token interface{}, approvalID uint, updatedApproval Core) (Core, error)
}

type ApprovalData interface {
	PostApproval(employeeID uint, newApproval Core) (Core, error)
	GetApproval() ([]Core, error)
	UpdateApproval(adminID uint, approvalID uint, updatedApproval Core) (Core, error)
}

package announcement

import "github.com/labstack/echo/v4"

type Core struct {
	ID      uint
	UserID  uint
	Type    string
	Title   string
	Message string
}

type AnnouncementHandler interface {
	PostAnnouncement() echo.HandlerFunc
	// GetAnnouncement() echo.HandlerFunc
	// DeleteAnnouncement() echo.HandlerFunc
}

type AnnouncementService interface {
	PostAnnouncement(token interface{}, newAnnouncement Core) (Core, error)
	// GetAnnouncement() ([]Core, error)
	// DeleteAnnouncement(token interface{}, announcementID string) error
}

type AnnouncementData interface {
	PostAnnouncement(adminID uint, newAnnouncement Core) (Core, error)
	// GetAnnouncement() ([]Core, error)
	// DeleteAnnouncement(adminID uint, announcementID uint) error
}

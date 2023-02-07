package announcement

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint
	Nip              string
	Name             string
	Type             string
	Title            string
	Message          string
	AnnouncementDate string
}

type AnnouncementHandler interface {
	PostAnnouncement() echo.HandlerFunc
	GetAnnouncement() echo.HandlerFunc
	GetAnnouncementDetail() echo.HandlerFunc
	DeleteAnnouncement() echo.HandlerFunc
}

type AnnouncementService interface {
	PostAnnouncement(token interface{}, newAnnouncement Core) (Core, error)
	GetAnnouncement() ([]Core, error)
	GetAnnouncementDetail(token interface{}, announcementID uint) (Core, error)
	DeleteAnnouncement(token interface{}, announcementID uint) error
}

type AnnouncementData interface {
	PostAnnouncement(adminID uint, newAnnouncement Core) (Core, error)
	GetAnnouncement() ([]Core, error)
	GetAnnouncementDetail(adminID uint, announcementID uint) (Core, error)
	DeleteAnnouncement(adminID uint, announcementID uint) error
}

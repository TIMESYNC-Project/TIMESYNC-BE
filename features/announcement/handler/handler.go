package handler

import (
	"log"
	"net/http"
	"timesync-be/features/announcement"
	"timesync-be/helper"

	"github.com/labstack/echo/v4"
)

type announcementControll struct {
	srv announcement.AnnouncementService
}

func New(as announcement.AnnouncementService) announcement.AnnouncementHandler {
	return &announcementControll{
		srv: as,
	}
}

func (ac *announcementControll) PostAnnouncement() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := PostAnnouncementRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}
		cnv := ReqToCore(input)
		res, err := ac.srv.PostAnnouncement(token, *cnv)

		if err != nil {
			log.Println("error post content : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "send announcement message to employee success", ToPostAnnouncementReponse(res)))
	}
}

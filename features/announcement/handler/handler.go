package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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
		input := PostAnnouncementRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}
		res, err := ac.srv.PostAnnouncement(c.Get("user"), *ReqToCore(input))

		if err != nil {
			log.Println("error post content : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "send announcement message to employee success", ToPostAnnouncementReponse(res)))
	}
}

func (ac *announcementControll) GetAnnouncement() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.srv.GetAnnouncement()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []ShowAllAnnouncement{}
		for _, val := range res {
			result = append(result, ShowAllAnnouncementJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success get all announcement",
		})

	}
}

func (ac *announcementControll) GetAnnouncementDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")

		announcementID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		res, err := ac.srv.GetAnnouncementDetail(token, uint(announcementID))

		if err != nil {
			if strings.Contains(err.Error(), "user") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "user not found",
				})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "announcement not found",
				})
			}
		}

		// result := []ShowAllAnnouncement{}
		// for _, val := range res {
		// 	result = append(result, ShowAllAnnouncementJson(val))
		// }
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "success get announcement details", ShowAllAnnouncementJson(res)))

	}
}

func (ac *announcementControll) EmployeeInbox() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := ac.srv.EmployeeInbox(token)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []ShowAllAnnouncement{}
		for _, val := range res {
			result = append(result, ShowAllAnnouncementJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show employee inbox message",
		})
	}
}

func (ac *announcementControll) DeleteAnnouncement() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")

		announcementID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		err = ac.srv.DeleteAnnouncement(token, uint(announcementID))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "success delete announcement")
	}
}

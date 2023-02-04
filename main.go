package main

import (
	"log"
	"timesync-be/config"
	announData "timesync-be/features/announcement/data"
	announHdl "timesync-be/features/announcement/handler"
	announSrv "timesync-be/features/announcement/services"
	attData "timesync-be/features/attendance/data"
	attHdl "timesync-be/features/attendance/handler"
	attSrv "timesync-be/features/attendance/services"
	usrData "timesync-be/features/user/data"
	usrHdl "timesync-be/features/user/handler"
	usrSrv "timesync-be/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)

	uData := usrData.New(db)
	uSrv := usrSrv.New(uData)
	uHdl := usrHdl.New(uSrv)
	attendData := attData.New(db)
	attendSrv := attSrv.New(attendData)
	attendHdl := attHdl.New(attendSrv)
	announcementData := announData.New(db)
	announcementSrv := announSrv.New(announcementData)
	announcementHdl := announHdl.New(announcementSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	//User
	e.POST("/register", uHdl.Register(), middleware.JWT([]byte(config.JWTKey)))
	e.POST("/login", uHdl.Login())
	e.DELETE("/employees/:id", uHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))
	e.POST("/register/csv", uHdl.Csv(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/employees/profile", uHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/employees/:id", uHdl.ProfileEmployee(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/employees/:id", uHdl.AdminEditEmployee(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/employees", uHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/employees", uHdl.GetAllEmployee())
	e.POST("/announcements", announcementHdl.PostAnnouncement(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/announcements", announcementHdl.PostAnnouncement(), middleware.JWT([]byte(config.JWTKey)))

	//attendance for emloyees
	e.GET("/attendance/location", attendHdl.GetLL())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

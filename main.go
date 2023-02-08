package main

import (
	"log"
	"timesync-be/config"
	announData "timesync-be/features/announcement/data"
	announHdl "timesync-be/features/announcement/handler"
	announSrv "timesync-be/features/announcement/services"
	approvalData "timesync-be/features/approval/data"
	approvalHdl "timesync-be/features/approval/handler"
	approvalSrv "timesync-be/features/approval/services"
	attData "timesync-be/features/attendance/data"
	attHdl "timesync-be/features/attendance/handler"
	attSrv "timesync-be/features/attendance/services"
	cmpData "timesync-be/features/company/data"
	cmpHdl "timesync-be/features/company/handler"
	cmpSrv "timesync-be/features/company/services"
	stData "timesync-be/features/setting/data"
	stHdl "timesync-be/features/setting/handler"
	stSrv "timesync-be/features/setting/services"
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
	setData := stData.New(db)
	setSrv := stSrv.New(setData)
	setHdl := stHdl.New(setSrv)
	approvalData := approvalData.New(db)
	approvalSrv := approvalSrv.New(approvalData)
	approvalHdl := approvalHdl.New(approvalSrv)
	cmData := cmpData.New(db)
	cmSrv := cmpSrv.New(cmData)
	cmHdl := cmpHdl.New(cmSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	//Users
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
	e.GET("/announcements", announcementHdl.GetAnnouncement(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/announcements/:id", announcementHdl.GetAnnouncementDetail(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/announcements/:id", announcementHdl.DeleteAnnouncement(), middleware.JWT([]byte(config.JWTKey)))
	e.POST("/approvals", approvalHdl.PostApproval(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/approvals", approvalHdl.GetApproval(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/approvals/:id", approvalHdl.UpdateApproval(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/approvals/:id", approvalHdl.ApprovalDetail())
	e.GET("/search", uHdl.Search(), middleware.JWT([]byte(config.JWTKey)))

	//attendances for emloyees
	e.GET("/inbox", announcementHdl.EmployeeInbox(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/employee/approvals", approvalHdl.EmployeeApprovalRecord(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/attendances/location", attendHdl.GetLL())
	e.POST("/attendances", attendHdl.ClockIn(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/attendances", attendHdl.ClockOut(), middleware.JWT([]byte(config.JWTKey)))
	e.POST("/attendances/:id", attendHdl.AttendanceFromAdmin(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/attendances", attendHdl.Record(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/presences", attendHdl.GetPresenceToday(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/presences/total", attendHdl.GetPresenceTotalToday(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/record/:id", attendHdl.RecordByID(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/presences/detail/:id", attendHdl.GetPresenceDetail(), middleware.JWT([]byte(config.JWTKey)))

	//setting
	e.GET("/setting", setHdl.GetSetting())
	e.PUT("/setting", setHdl.EditSetting(), middleware.JWT([]byte(config.JWTKey)))

	//company
	e.GET("/companies", cmHdl.GetProfile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/companies", cmHdl.EditProfile(), middleware.JWT([]byte(config.JWTKey)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

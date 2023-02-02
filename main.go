package main

import (
	"timesync-be/config"
	usrData "timesync-be/features/user/data"
	usrHdl "timesync-be/features/user/handler"
	usrSrv "timesync-be/features/user/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)

	uData := usrData.New(db)
	uSrv := usrSrv.New(uData)
	uHdl := usrHdl.New(uSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
}

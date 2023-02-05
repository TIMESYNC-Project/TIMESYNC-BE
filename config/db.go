package config

import (
	"fmt"
	"log"
	aData "timesync-be/features/announcement/data"
	uData "timesync-be/features/user/data"

	attendData "timesync-be/features/attendance/data"
	settData "timesync-be/features/setting/data"
	usrData "timesync-be/features/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dc DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.DBUser, dc.DBPass, dc.DBHost, dc.DBPort, dc.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(uData.User{})
	db.AutoMigrate(aData.Announcement{})
	db.AutoMigrate(usrData.User{})
	db.AutoMigrate(attendData.Attendance{})
	db.AutoMigrate(settData.Setting{})
}

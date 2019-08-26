package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GormDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "2u4!g3evS+"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "ecmonster"
	TIME := "?parseTime=true"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME + TIME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	return db
}

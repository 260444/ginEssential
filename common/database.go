package common

import (
	"fmt"

	"github.com/260444/ginEssential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	// driveName := "mysql"
	user := "root"
	pass := "123456"
	host := "localhost"
	port := "3306"
	database := "ceshi"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	DB = db
}

func GetDB() *gorm.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}

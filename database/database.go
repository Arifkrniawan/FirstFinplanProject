package database

import (
	"github.com/Arifkrniawan/FirstFinplanProject/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnection() {

	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/databasehobby"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Hobby{}, &models.MapUserHobby{})

	DB = db
}

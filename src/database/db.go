package database

import (
	"go-ambassador/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Could not connect with the dataabse!")
	}

}

func AutoMigrate() {
	DB.AutoMigrate(
		models.User{},
		models.Product{},
		models.Link{},
		models.Order{},
		models.OrderItem{},
	)
}

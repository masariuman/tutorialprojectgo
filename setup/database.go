package setup

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connect *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/tutorialproject?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// database.AutoMigrate(Artikel{})

	Connect = database
}

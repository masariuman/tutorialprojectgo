package main

import (
	"log"
	"tutorialProject/controller"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/tutorialproject?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi Error")
	} else {
		log.Fatal("Koneksi Sukses")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", controller.RootHandler)
	v1.GET("/halo", controller.HaloHandler)
	v1.GET("/data/:id", controller.DataHandler)
	v1.GET("/query", controller.QueryHandler)
	v1.GET("/data/:id/:judul", controller.DataLebihVarHandler)
	v1.POST("/data", controller.PostDataHandler)

	router.Run()

}

package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"tutorialProject/struct/data"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama": "Arif Setiawan",
		"bio":  "A Software Engineer",
	})
}

func HaloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"judul": "SELAMAT !",
		"isi":   "Selamat ! Anda masuk ke url halo !",
	})
}

func DataHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func QueryHandler(c *gin.Context) {
	judul := c.Query("judul")
	isi := c.Query("isi")
	c.JSON(http.StatusOK, gin.H{"judul": judul, "isi": isi})
}

func DataLebihVarHandler(c *gin.Context) {
	id := c.Param("id")
	judul := c.Param("judul")
	c.JSON(http.StatusOK, gin.H{"id": id, "judul": judul})
}

func PostDataHandler(c *gin.Context) {
	var artikel data.Artikel
	err := c.ShouldBindJSON(&artikel)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error pada %s, Kondisi : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	err = data.Connect.Create(&artikel).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Disimpan."})
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"tutorialProject/repository"
	"tutorialProject/setup"
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

	err = setup.Connect.Create(&artikel).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Disimpan."})
}

func ReadHandler(c *gin.Context) {

	articleRepository := repository.NewArticleRepository(setup.Connect)
	artikels, err := articleRepository.FindAll()

	if err != nil {
		panic(err)
	}

	for _, Q := range artikels { //ini seperti foreach
		fmt.Println("judul : ", Q.Judul)
		fmt.Println("isi : ", Q.Isi)
	}

	c.JSON(http.StatusOK, artikels)
}

func ReadidHandler(c *gin.Context) {
	var artikel data.Artikel
	id := c.Param("id")

	err := setup.Connect.Debug().First(&artikel, id).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, artikel)
}

func UpdateHandler(c *gin.Context) {
	var artikel data.Artikel
	id := c.Param("id")

	err := c.ShouldBindJSON(&artikel)
	if err != nil {
		panic(err)
	}

	err = setup.Connect.Model(&artikel).Where("id = ?", id).Updates(&artikel).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Diubah."})
}

func DeleteHandler(c *gin.Context) {
	var artikel data.Artikel

	var input struct {
		Id json.Number
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	id, _ := input.Id.Int64()

	err = setup.Connect.Debug().Delete(&artikel, id).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Dihapus."})
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tutorialProject/repository"
	"tutorialProject/service"
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
	var artikel data.ArtikelRequest
	err := c.ShouldBindJSON(&artikel)

	// articleRepository := repository.NewArticleRepository(setup.Connect)
	articleService := service.NewArticleService(repository.NewArticleRepository(setup.Connect))
	_, err = articleService.Store(artikel, err, c)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Disimpan."})
}

func ReadHandler(c *gin.Context) {

	// artikels, err := articleRepository.FindAll() // cara langsung tembak repo
	articleService := service.NewArticleService(repository.NewArticleRepository(setup.Connect))
	articles, err := articleService.FindAll()

	if err != nil {
		panic(err)
	}

	// for _, Q := range artikels { //ini seperti foreach
	// 	fmt.Println("judul : ", Q.Judul)
	// 	fmt.Println("isi : ", Q.Isi)
	// }

	c.JSON(http.StatusOK, articles)
}

func ReadidHandler(c *gin.Context) {
	id := c.Param("id")
	// articleRepository := repository.NewArticleRepository(setup.Connect)
	articleService := service.NewArticleService(repository.NewArticleRepository(setup.Connect))
	artikel, err := articleService.FindById(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, artikel)
}

func UpdateHandler(c *gin.Context) {
	var artikel data.ArtikelRequest
	id := c.Param("id")

	err := c.ShouldBindJSON(&artikel)
	if err != nil {
		panic(err)
	}

	// articleRepository := repository.NewArticleRepository(setup.Connect)
	articleService := service.NewArticleService(repository.NewArticleRepository(setup.Connect))
	_, err = articleService.Update(id, artikel)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Diubah."})
}

func DeleteHandler(c *gin.Context) {
	var input struct {
		Id string
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	// articleRepository := repository.NewArticleRepository(setup.Connect)
	articleService := service.NewArticleService(repository.NewArticleRepository(setup.Connect))
	_, err = articleService.Delete(input.Id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Dihapus."})
}

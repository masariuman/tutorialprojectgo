package service

import (
	"fmt"
	"net/http"
	"tutorialProject/repository"
	"tutorialProject/struct/data"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArticleService interface {
	FindAll() ([]data.Artikel, error)
	FindById(ID string) (data.Artikel, error)
	Store(artikelRequest data.ArtikelRequest, err error, c *gin.Context) (data.Artikel, error)
	Update(ID string, artikel data.ArtikelRequest) (data.Artikel, error)
	Delete(ID string) (data.Artikel, error)
}

type service struct {
	repository repository.ArticleRepository
}

func NewArticleService(repository repository.ArticleRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]data.Artikel, error) {
	// articles, err := s.repository.FindAll() //cara repota
	// return articles, err
	return s.repository.FindAll() //cara singkat
}

func (s *service) FindById(ID string) (data.Artikel, error) {
	return s.repository.FindById(ID)
}

func (s *service) Store(artikelRequest data.ArtikelRequest, err error, c *gin.Context) (data.Artikel, error) {
	if err != nil {
		var article data.Artikel
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error pada %s, Kondisi : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return article, err
	}
	article := data.Artikel{
		Judul:     artikelRequest.Judul,
		Deskripsi: artikelRequest.Deskripsi,
		Isi:       artikelRequest.Isi,
		Penulis:   artikelRequest.Penulis,
	}
	return s.repository.Store(article)
}

func (s *service) Update(ID string, artikelRequest data.ArtikelRequest) (data.Artikel, error) {
	article := data.Artikel{
		Judul:     artikelRequest.Judul,
		Deskripsi: artikelRequest.Deskripsi,
		Isi:       artikelRequest.Isi,
		Penulis:   artikelRequest.Penulis,
	}
	return s.repository.Update(ID, article)
}

func (s *service) Delete(ID string) (data.Artikel, error) {
	return s.repository.Delete(ID)
}

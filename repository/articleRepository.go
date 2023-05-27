package repository

import (
	"tutorialProject/struct/data"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll() ([]data.Artikel, error)
	FindById(ID string) (data.Artikel, error)
	Store(artikel data.Artikel) (data.Artikel, error)
	Update(ID string, artikel data.Artikel) (data.Artikel, error)
	Delete(ID string) (data.Artikel, error)
}

type repository struct {
	database *gorm.DB
}

func NewArticleRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAll() ([]data.Artikel, error) {
	var artikels []data.Artikel
	err := r.database.Debug().Find(&artikels).Error
	return artikels, err
}

func (r *repository) FindById(ID string) (data.Artikel, error) {
	var artikel data.Artikel
	err := r.database.Debug().First(&artikel, ID).Error
	return artikel, err
}

func (r *repository) Store(artikel data.Artikel) (data.Artikel, error) {
	err := r.database.Create(&artikel).Error
	return artikel, err
}

func (r *repository) Update(ID string, artikel data.Artikel) (data.Artikel, error) {
	// var artikel data.Artikel
	err := r.database.Model(&artikel).Where("id = ?", ID).Updates(&artikel).Error
	return artikel, err
}

func (r *repository) Delete(ID string) (data.Artikel, error) {
	var artikel data.Artikel
	err := r.database.Debug().Delete(&artikel, ID).Error
	return artikel, err
}

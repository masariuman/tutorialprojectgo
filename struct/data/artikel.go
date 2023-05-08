package data

import "time"

type Artikel struct {
	Id        uint      `gorm:"primaryKey"`
	Judul     string    `gorm:"type:varchar(255)" 	json:"judul"`
	Deskripsi string    `gorm:"type:text" 			json:"deskripsi"`
	Isi       string    `gorm:"type:text" 			json:"isi"`
	Penulis   string    `gorm:"type:varchar(255)" 	json:"penulis"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}

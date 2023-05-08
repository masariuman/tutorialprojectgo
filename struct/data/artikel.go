package data

import "time"

type Artikel struct {
	Id        uint `gorm:"primaryKey" json:"id"`
	Judul     string
	Deskripsi string
	Isi       string
	Penulis   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

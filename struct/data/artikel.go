package data

import "time"

type Artikel struct {
	ID        uint
	Judul     string
	Deskripsi string
	Isi       string
	Penulis   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

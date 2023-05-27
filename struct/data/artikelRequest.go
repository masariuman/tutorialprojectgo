package data

type ArtikelRequest struct {
	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi"`
	Isi       string `json:"isi"`
	Penulis   string `json:"penulis"`
}

//a

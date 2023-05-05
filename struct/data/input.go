package data

type DataInput struct {
	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
}

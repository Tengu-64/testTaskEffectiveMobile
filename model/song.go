package model

type Song struct {
	Id          uint `gorm:"primaryKey"`
	Group       string
	Name        string
	ReleaseDate string
	Text        string
	Link        string
}

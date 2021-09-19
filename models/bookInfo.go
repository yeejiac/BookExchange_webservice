package models

import "time"

type BookCatelog int

const (
	North BookCatelog = iota
	Action_and_Adventure
	Horror
	Fantasy
	Comic_Book
	Literary_Fiction
	Romance
	SciFi
	Other
)

type BookInfo struct {
	BookName    string      `json:"bookname"`
	ISBN        string      `json:"isbn"`
	Writer      string      `json:"writer"`
	Publisher   string      `json:"publisher"`
	PublishDate time.Time   `json:"publishdate"`
	BookCatelog BookCatelog `json:"bookCatelog"`
}

package models

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
	BookCatelog BookCatelog `json:"bookcatelog"`
	ImageName   string      `json:"imagename"`
}

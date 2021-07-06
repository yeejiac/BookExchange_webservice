package models

import "time"

type BookCatelog struct {
	ID string `json:"id"`
}

type BookInfo struct {
	BookName    string    `json:"bookname"`
	ISBN        string    `json:"isbn"`
	Writer      string    `json:"writer"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publishdate"`
}

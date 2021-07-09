package models

type User struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Validation bool   `json:"vaidation"`
}

type Privacy int

const (
	PUBLIC Privacy = iota
	PRIVATE
	PROTECT
)

package models

type User struct {
	Username          string        `json:"username"`
	Name              string        `json:"name"`
	Age               int           `json:"age"`
	Email             string        `json:"email"`
	Password          string        `json:"password"`
	Validation        bool          `json:"vaidation"`
	Address           string        `json:"address"`
	TotalExchangeTime int           `json:"totalexchangetime"`
	KKC               int           `json:"kkc"`
	Status            AccountStatus `json:"Status"`
}

type AccountStatus int

const (
	ENABLE = iota
	DISABLE
)

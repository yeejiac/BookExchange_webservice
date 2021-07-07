package models

type UserAccount struct {
	Username          string        `json:"username"`
	TotalExchangeTime int           `json:"totalexchangetime"`
	KKC               int           `json:"kkc"`
	Status            AccountStatus `json:"Status"`
}

type AccountStatus int

const (
	ABLE = iota
	DISABLE
)

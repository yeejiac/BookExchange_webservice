package models

type ExchangeInfo struct {
	Username string   `json:"username"`
	Name     string   `json:"name"`
	GroupID  string   `json:"groupid"`
	Address  string   `json:"address"`
	Book     BookInfo `json:"book"`
}

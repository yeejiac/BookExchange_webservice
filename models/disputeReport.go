package models

type Dispute struct {
	Reporter string `json:"reporter"`
	Accuse   string `json:"accuse"`
	Issue    string `json:"issue"`
}

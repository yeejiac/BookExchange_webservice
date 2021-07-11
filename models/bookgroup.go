package models

import "time"

type ExchangeGroup struct {
	GroupName        string      `json:"groupname"`
	GroupID          string      `json:"groupid"`
	Holder           string      `json:"holder"`
	CateLog          BookCatelog `json:"catelog"`
	Privacy          Privacy     `json:"privacy"`
	EstablishTime    time.Time   `json:"establishtime"`
	ExpireTime       time.Time   `json:"expiretime"`
	GroupSize        int         `json:"groupsize"`
	GroupDescription string      `json:"groupdescription"`
	Attender         []User      `json:"attender"`
}

type Privacy int

const (
	PUBLIC Privacy = iota
	PRIVATE
	PROTECT
)

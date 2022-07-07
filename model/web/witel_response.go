package web

import "time"

type WitelResponse struct {
	Id       string           `json:"id"`
	Alias    string           `json:"alias"`
	Nama     string           `json:"nama"`
	Regional RegionalResponse `json:"regional"`
	Stos     []StoResponse    `json:"stos"`
	CreateAt time.Time        `json:"create_at"`
	UpdateAt time.Time        `json:"update_at"`
}

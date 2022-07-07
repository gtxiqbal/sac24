package web

import "time"

type StoResponse struct {
	Id       string        `json:"id"`
	Alias    string        `json:"alias"`
	Nama     string        `json:"nama"`
	Witel    WitelResponse `json:"witel"`
	CreateAt time.Time     `json:"create_at"`
	UpdateAt time.Time     `json:"update_at"`
}

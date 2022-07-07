package domain

import "time"

type Sto struct {
	Id       string
	Alias    string
	Nama     string
	Witel    Witel
	Gpon     []Gpon
	CreateAt time.Time
	UpdateAt time.Time
}

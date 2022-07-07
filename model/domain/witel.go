package domain

import "time"

type Witel struct {
	Id       string
	Alias    string
	Nama     string
	Regional Regional
	Stos     []Sto
	CreateAt time.Time
	UpdateAt time.Time
}

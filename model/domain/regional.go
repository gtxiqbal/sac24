package domain

import (
	"time"
)

type Regional struct {
	Id       string
	Nama     string
	witels   []Witel
	CreateAt time.Time
	UpdateAt time.Time
}

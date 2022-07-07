package web

import (
	"time"
)

type RegionalResponse struct {
	Id       string          `json:"id"`
	Nama     string          `json:"nama"`
	Witels   []WitelResponse `json:"witels"`
	CreateAt time.Time       `json:"create_at"`
	UpdateAt time.Time       `json:"update_at"`
}

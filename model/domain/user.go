package domain

import "time"

type User struct {
	Username            string
	FirstName           string
	LastName            string
	Password            string
	Email               string
	Enabled             bool
	ExpiredDate         time.Time
	ExpiredPasswordDate time.Time
	TelegramId          int
	DefaultUser         bool
	FailPass            int
	LimitLogin          int
	LimitFail           int
	CreateAt            time.Time
	UpdateAt            time.Time
}

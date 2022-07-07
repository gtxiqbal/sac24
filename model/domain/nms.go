package domain

import "time"

type Nms struct {
	Id       string
	Nama     string
	Vendor   string
	IpServer string
	PortTl1  int32
	Protocol string
	Username string
	Password string
	Gpon     []Gpon
	CreateAt time.Time
	UpdateAt time.Time
}

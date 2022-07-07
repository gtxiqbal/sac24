package domain

import "time"

type Gpon struct {
	Id             string
	Hostname       string
	IpAddress      string
	VlanInet       int32
	VlanVoice      string
	DefaultUserNms bool
	Username       string
	Password       string
	Nms            Nms
	Sto            Sto
	CreateAt       time.Time
	UpdateAt       time.Time
}

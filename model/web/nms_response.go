package web

import "time"

type NmsResponse struct {
	Id       string         `json:"id"`
	Nama     string         `json:"nama"`
	Vendor   string         `json:"vendor"`
	IpServer string         `json:"ip_server"`
	PortTl1  int32          `json:"port_tl1"`
	Protocol string         `json:"protocol"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Gpons    []GponResponse `json:"gpons"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at"`
}

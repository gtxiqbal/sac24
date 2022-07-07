package web

type NmsRequest struct {
	Id       string `json:"id"`
	Nama     string `json:"nama"`
	Vendor   string `json:"vendor"`
	IpServer string `json:"ip_server"`
	PortTl1  int32  `json:"port_tl1"`
	Protocol string `json:"protocol"`
	Username string `json:"username"`
	Password string `json:"password"`
}

package web

type GponRequest struct {
	Id             string      `json:"id"`
	Hostname       string      `json:"hostname"`
	IpAddress      string      `json:"ip_address"`
	VlanInet       int32       `json:"vlan_inet"`
	VlanVoice      string      `json:"vlan_voice"`
	DefaultUserNms bool        `json:"default_user_nms"`
	Username       string      `json:"username"`
	Password       string      `json:"password"`
	Nms            NmsResponse `json:"nms"`
	Sto            StoResponse `json:"sto"`
}

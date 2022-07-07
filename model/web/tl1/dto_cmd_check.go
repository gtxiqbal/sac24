package tl1

import "github.com/gtxiqbal/sac24/model/web"

type DtoCmdCheck struct {
	Nms     web.NmsRequest `json:"nms"`
	Login   string         `json:"login"`
	Logout  string         `json:"logout"`
	CmdType string         `json:"cmd_type"`
	CmdList DtoCmdDetail   `json:"cmd_list"`
	Gpons   []DtoConfig    `json:"gpons"`
	IpGpon  string         `json:"ip_gpon"`
}

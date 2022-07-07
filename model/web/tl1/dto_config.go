package tl1

import "github.com/gtxiqbal/sac24/model/web"

type DtoConfig struct {
	IpGpon         string         `json:"ip_gpon"`
	SlotPort       string         `json:"slot_port"`
	OnuId          string         `json:"onu_id"`
	OnuType        string         `json:"onu_type"`
	OnuName        string         `json:"onu_name"`
	OpenOnt        string         `json:"open_ont"`
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	Status         int            `json:"status"`
	DefaultUserNms bool           `json:"default_user_nms"`
	Nms            web.NmsRequest `json:"nms"`
}

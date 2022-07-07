package tl1

import "github.com/gtxiqbal/sac24/model/web/tl1"

type CommandCheck interface {
	CheckUnReg(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck
	CheckReg(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck
	CheckService(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck
}

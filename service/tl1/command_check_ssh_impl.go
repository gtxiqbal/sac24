package tl1

import (
	"encoding/base64"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/model/web/tl1"
	"strings"
)

type CommandCheckSSH CommandCheck

type CommandCheckSSHImpl struct {
}

func NewCommandCheckSSHImpl() CommandCheckSSH {
	return &CommandCheckSSHImpl{}
}

func (sshService *CommandCheckSSHImpl) CheckUnReg(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck {
	return sshService.check(dtoConfigs, "UNREG")
}

func (sshService *CommandCheckSSHImpl) CheckReg(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck {
	return sshService.check(dtoConfigs, "REG")
}

func (sshService *CommandCheckSSHImpl) CheckService(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck {
	return sshService.check(dtoConfigs, "SERVICE")
}

func (sshService *CommandCheckSSHImpl) check(dtoConfigs []tl1.DtoConfig, cmdType string) []tl1.DtoCmdCheck {
	var cmdChecks []tl1.DtoCmdCheck
	var ipServers []web.NmsRequest
	for _, config := range dtoConfigs {
		request := config.Nms
		ipServers = append(ipServers, request)
	}
	ipServers = helper.SliceUniqueNmsRequest(ipServers)

	for _, nms := range ipServers {
		var cmdCheck tl1.DtoCmdCheck
		var gpons []tl1.DtoConfig

		for _, config := range dtoConfigs {
			if strings.EqualFold(nms.IpServer, config.Nms.IpServer) {
				gpon := tl1.DtoConfig{
					IpGpon:   config.IpGpon,
					SlotPort: config.SlotPort,
					OnuId:    config.OnuId,
					Username: config.Username,
					Password: config.Password,
				}
				username, err := base64.StdEncoding.DecodeString(nms.Username)
				helper.PanicIfError(err)
				password, err := base64.StdEncoding.DecodeString(nms.Password)
				helper.PanicIfError(err)
				if !gpon.DefaultUserNms {
					if gpon.Username != "" && gpon.Password != "" {
						username, err = base64.StdEncoding.DecodeString(gpon.Username)
						helper.PanicIfError(err)
						password, err = base64.StdEncoding.DecodeString(gpon.Password)
						helper.PanicIfError(err)
					}
				}
				gpon.Username = string(username)
				gpon.Password = string(password)
				gpons = append(gpons, gpon)
			}
		}
		cmdCheck.Nms = nms
		username, err := base64.StdEncoding.DecodeString(nms.Username)
		helper.PanicIfError(err)
		password, err := base64.StdEncoding.DecodeString(nms.Password)
		helper.PanicIfError(err)
		cmdCheck.Nms.Username = string(username)
		cmdCheck.Nms.Password = string(password)

		cmdCheck.Gpons = gpons
		cmdChecks = append(cmdChecks, cmdCheck)
	}

	return cmdChecks
}

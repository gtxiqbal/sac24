package tl1

import (
	"encoding/base64"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/model/web/tl1"
	"strings"
)

type CommandCheckTl1 CommandCheck

type CommandCheckTL1Impl struct {
	ConfigUnRegTimeSleep   int
	ConfigRegTimeSleep     int
	ConfigServiceTimeSleep int
}

func NewCommandCheck(configUnRegTimeSleep int, configRegTimeSleep int, configServiceTimeSleep int) CommandCheckTl1 {
	return &CommandCheckTL1Impl{ConfigUnRegTimeSleep: configUnRegTimeSleep, ConfigRegTimeSleep: configRegTimeSleep, ConfigServiceTimeSleep: configServiceTimeSleep}
}

func (tl1Service *CommandCheckTL1Impl) CheckUnReg(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck {
	return tl1Service.check(dtoConfigs, "UNREG")
}

func (tl1Service *CommandCheckTL1Impl) CheckReg(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck {
	return tl1Service.check(dtoConfigs, "REG")
}

func (tl1Service *CommandCheckTL1Impl) CheckService(dtoConfigs []tl1.DtoConfig) []tl1.DtoCmdCheck {
	return tl1Service.check(dtoConfigs, "SERVICE")
}

func (tl1Service *CommandCheckTL1Impl) check(dtoConfigs []tl1.DtoConfig, cmdType string) []tl1.DtoCmdCheck {
	var cmdChecks []tl1.DtoCmdCheck
	var ipServers []web.NmsRequest
	for _, config := range dtoConfigs {
		request := config.Nms
		ipServers = append(ipServers, request)
	}
	ipServers = helper.SliceUniqueNmsRequest(ipServers)

	desc := "CEK UNREGISTER ONT"
	if strings.EqualFold("REG", cmdType) {
		desc = "CEK REGISTER ONT"
	} else if strings.EqualFold("SERVICE", cmdType) {
		desc = "CEK SERVICE ONT"
	}

	for _, nms := range ipServers {
		var cmdCheck tl1.DtoCmdCheck
		var cmdList []string
		vendor := nms.Vendor
		onuIdType := "MAC"
		if !strings.EqualFold("ZTE", vendor) {
			onuIdType = "SN"
		}

		for _, config := range dtoConfigs {
			if strings.EqualFold(nms.IpServer, config.Nms.IpServer) {
				ipGpon := config.IpGpon
				var cmd string
				if strings.EqualFold("UNREG", cmdType) {
					cmd = "LST-UNREGONU::OLTID=" + ipGpon + ":::;"
					cmdList = append(cmdList, cmd)
				} else if strings.EqualFold("REG", cmdType) {
					cmd = "LST-REGONU::OLTID=" + ipGpon
					if !strings.EqualFold("", config.SlotPort) {
						cmd += ",PONID=1-1-" + config.SlotPort
					}
					if !strings.EqualFold("", config.OnuId) {
						cmd += ",ONUIDTYPE=" + onuIdType + ",ONUID=" + config.OnuId
					}
					cmd += ":::;"
					cmdList = append(cmdList, cmd)
				} else if strings.EqualFold("SERVICE", cmdType) {
					slotPort := config.SlotPort
					onuId := config.OnuId

					if strings.EqualFold("ALU", vendor) || strings.EqualFold("FH", vendor) || strings.EqualFold("ZTE", vendor) {
						cmd = "LST-ONU::OLTID=" + ipGpon + ",PONID=1-1-" + slotPort + ",ONUIDTYPE=" + onuIdType + ",ONUID=" + onuId + ":::;"
						cmdList = append(cmdList, cmd)

						cmd = "LST-ONUSTATE::OLTID=" + ipGpon + ",PONID=1-1-" + slotPort + ",ONUIDTYPE=" + onuIdType + ",ONUID=" + onuId + ":::;"
						cmdList = append(cmdList, cmd)

						if strings.EqualFold("FH", vendor) {
							cmd = "LST-ONUWANSERVICECFG::OLTID=" + ipGpon + ",PONID=1-1-" + slotPort + ",ONUIDTYPE=" + onuIdType + ",ONUID=" + onuId + ":::;"
							cmdList = append(cmdList, cmd)

							cmd = "LST-POTS::OLTID=" + ipGpon + ",PONID=1-1-" + slotPort + ",ONUIDTYPE=MAC,ONUID=" + onuId + ":::;"
							cmdList = append(cmdList, cmd)
						} else if strings.EqualFold("FH", vendor) {
							cmd = "LST-SERVICEPORT::DID=" + ipGpon + ",OID=" + onuId + ":::;"
							cmdList = append(cmdList, cmd)

							cmd = "LST-ONUWANIP::DID=" + ipGpon + ",OID=" + onuId + ":::;"
							cmdList = append(cmdList, cmd)
						}

						cmd = "LST-POTSINFO::OLTID=" + ipGpon + ",PONID=1-1-" + slotPort + ",ONUIDTYPE=" + onuIdType + ",ONUID=" + onuId + ":::;"
						cmdList = append(cmdList, cmd)

						cmd = "LST-PORTVLAN::OLTID=" + ipGpon + ",PONID=1-1-" + slotPort + ",ONUIDTYPE=" + onuIdType + ",ONUID=" + onuId + ":::;"
						cmdList = append(cmdList, cmd)
					}
				}
			}
		}

		cmdCheck.Nms = nms
		username, err := base64.StdEncoding.DecodeString(nms.Username)
		helper.PanicIfError(err)
		password, err := base64.StdEncoding.DecodeString(nms.Password)
		helper.PanicIfError(err)

		if !strings.EqualFold("ZTE", vendor) {
			cmdCheck.Login = "LOGIN:::CTAG::UN=" + string(username) + ",PWD=" + string(password) + ";"
			cmdCheck.Logout = "LOGOUT:::CTAG::;"
		}

		timeSleep := tl1Service.ConfigUnRegTimeSleep
		if strings.EqualFold("REG", cmdType) {
			timeSleep = tl1Service.ConfigRegTimeSleep
		} else if strings.EqualFold("SERVICE", cmdType) {
			timeSleep = tl1Service.ConfigServiceTimeSleep
		}
		cmdCheck.CmdList = tl1.DtoCmdDetail{Desc: desc, Cmds: cmdList, TimeSleep: timeSleep}
		cmdChecks = append(cmdChecks, cmdCheck)
	}
	return cmdChecks
}

package service

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/model/web/telegram/request"
	tl12 "github.com/gtxiqbal/sac24/model/web/tl1"
	"github.com/gtxiqbal/sac24/repository"
	"github.com/gtxiqbal/sac24/service/tl1"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ConfigViaTelegramServiceImpl struct {
	*sql.DB
	repository.GponRepository
	repository.UserRepository
	tl1.CommandCheckSSH
	tl1.CommandCheckTl1
	BotToken string
}

func NewConfigViaTelegramServiceImpl(DB *sql.DB, gponRepository repository.GponRepository, userRepository repository.UserRepository, commandCheckSSH tl1.CommandCheckSSH, commandCheckTl1 tl1.CommandCheckTl1, botToken string) ConfigViaTelegramService {
	return &ConfigViaTelegramServiceImpl{DB: DB, GponRepository: gponRepository, UserRepository: userRepository, CommandCheckSSH: commandCheckSSH, CommandCheckTl1: commandCheckTl1, BotToken: botToken}
}

func (service *ConfigViaTelegramServiceImpl) checkTelegramId(ctx context.Context, request request.SendMessageRequest) bool {
	now := time.Now().In(time.Local)
	dateExp := time.Date(2022, 07, 31, 23, 59, 59, 0000000, time.Local)

	if now.After(dateExp) {
		msg := "License Apliakasi Sudah Expired, Silahkan Hubungi Admin <b>@p_fjr</b>"
		err := errors.New(msg)
		service.doSendIfError(request, msg, err)
		return false
	}

	db := service.DB
	users, err := service.UserRepository.FindByTelegramId(ctx, db, request.ChatId)
	msg := "Akun Telegram Anda Tidak Terdaftar, Silahkan Hubungi Admin <b>@p_fjr</b>"
	isErr := service.doSendIfError(request, msg, err)
	if isErr {
		return false
	}

	user := users[0]
	if !user.Enabled {
		msg = "Akun Telegram Anda Tidak Aktif, Silahkan Hubungi Admin <b>@p_fjr</b>"
		err = errors.New(msg)
		service.doSendIfError(request, msg, err)
		return false
	}
	if now.After(user.ExpiredDate) {
		msg = "Akun Telegram Anda Sudah Expired, Silahkan Hubungi Admin <b>@p_fjr</b>"
		err = errors.New(msg)
		service.doSendIfError(request, msg, err)
		return false
	}
	return true
}

func (service *ConfigViaTelegramServiceImpl) DoSend(requestReply request.SendMessageRequest, dataBot string) {
	sizeDataBot := len(dataBot)
	if sizeDataBot > 4096 {
		maxSize := 4096
		amountSlice := sizeDataBot / maxSize
		start := 0
		end := maxSize
		for i := 0; i <= amountSlice; i++ {
			requestReply.Text = dataBot[start:end]
			start = start + maxSize
			end = end + maxSize
			if end > len(dataBot) {
				end = len(dataBot)
			}

			jsonByte, err := json.Marshal(&requestReply)
			helper.PanicIfError(err)

			_, _ = http.Post(
				"https://api.telegram.org/bot"+service.BotToken+"/sendMessage",
				"application/json",
				bytes.NewBuffer(jsonByte),
			)
		}
	} else {
		requestReply.Text = dataBot
		jsonByte, err := json.Marshal(&requestReply)
		helper.PanicIfError(err)

		_, _ = http.Post(
			"https://api.telegram.org/bot"+service.BotToken+"/sendMessage",
			"application/json",
			bytes.NewBuffer(jsonByte),
		)
		//helper.PanicIfError(err)
		//defer func(Body io.ReadCloser) {
		//	_ = Body.Close()
		//}(post.Body)

		//resp, err := ioutil.ReadAll(post.Body)
		//helper.PanicIfError(err)

		//var mapJson request.ResultResponse
		//_ = json.Unmarshal(resp, &mapJson)
		//helper.PanicIfError(err)
		//if mapJson.Ok {
		//	fmt.Println(string(resp))
		//} else {
		//	fmt.Println(mapJson.Description)
		//}
	}
}

func (service *ConfigViaTelegramServiceImpl) doSendIfError(request request.SendMessageRequest, msg string, err error) bool {
	if err != nil {
		dataBot := msg
		service.DoSend(request, dataBot)
		helper.LogIfError(err)
		return true
	}
	return false
}

func (service *ConfigViaTelegramServiceImpl) checkGpon(ctx context.Context, request request.SendMessageRequest, ipAddresses []string) (string, []domain.Gpon, bool) {
	db := service.DB
	protocol, err := service.GponRepository.FindProtocolByIpAddress(ctx, db, ipAddresses[0])
	msg := "Data GPON tidak ditemukan"
	isErr := service.doSendIfError(request, msg, err)
	if isErr {
		return "", nil, false
	}

	gpons, err := service.GponRepository.FindByIpAddressInAndProtocol(ctx, db, ipAddresses, protocol)
	msg = "Data GPON tidak ditemukan"
	isErr = service.doSendIfError(request, msg, err)
	if isErr {
		return "", nil, false
	}
	return protocol, gpons, true
}

func (service *ConfigViaTelegramServiceImpl) doCommandTl1(cmd string, conn net.Conn, reply []byte, results *[]string, requestReply request.SendMessageRequest) {
	_, err := conn.Write([]byte(cmd))
	msg := "Gagal Kirim Command"
	isErr := service.doSendIfError(requestReply, msg, err)
	if isErr {
		return
	}

	if !strings.HasPrefix(cmd, "LOGIN:::") && !strings.HasPrefix(cmd, "LOGOUT:::") {
		*results = append(*results, cmd)
		fmt.Println(cmd)
	}

	_, err = conn.Read(reply)
	msg = "Gagal Mendapatkan Hasil Konfigurasi"
	isErr = service.doSendIfError(requestReply, msg, err)
	if isErr {
		return
	}
	if !strings.HasPrefix(cmd, "LOGIN:::") && !strings.HasPrefix(cmd, "LOGOUT:::") {
		result := strings.ReplaceAll(string(bytes.Trim(reply, "\x00")), "\r", "")
		result = strings.ReplaceAll(result, "\n\n\n", "\n")
		result = strings.ReplaceAll(result, ";--------------------------------------------------------", "")
		result = strings.ReplaceAll(result, ";;", ";")
		*results = append(*results, result)
		fmt.Println(result)
	}
}

func (service *ConfigViaTelegramServiceImpl) doCloseTl1(cmd string, conn net.Conn, vendor string, results *[]string, requestReply request.SendMessageRequest) {
	defer func(conn net.Conn) {
		err := conn.Close()
		service.doSendIfError(requestReply, "Gagal Close Koneksi", err)
		fmt.Println("BERHASIL CLOSE KONEKSI")
	}(conn)

	if strings.ToUpper(vendor) != strings.ToUpper("ZTE") {
		service.doCommandTl1(cmd, conn, make([]byte, 10240), results, requestReply)
	}
}

func (service *ConfigViaTelegramServiceImpl) doProcessCheck(protocol string, configList []tl12.DtoConfig, dataBot string, request request.SendMessageRequest, typeConfig string) {
	defer helper.RecoverIfError()
	var results []string
	isErr := false
	if strings.ToUpper(protocol) == strings.ToUpper("SSH2") {
		command := `enable
scroll

display ont autofind all
quit
y


`
		cmdChecks := service.CommandCheckSSH.CheckUnReg(configList)
		for _, cmdCheck := range cmdChecks {
			service.doCommandSsh(cmdCheck, command, &results)

		}
		dataBot = strings.Join(results, "")
		request.ParseMode = ""
		service.DoSend(request, dataBot)
	} else {
		cmdChecks := service.CommandCheckTl1.CheckUnReg(configList)
		if typeConfig == "REG" {
			cmdChecks = service.CommandCheckTl1.CheckReg(configList)
		} else if typeConfig == "SERVICE" {
			cmdChecks = service.CommandCheckTl1.CheckService(configList)
		}
		for _, cmdCheck := range cmdChecks {
			conn, err := net.Dial("tcp", cmdCheck.Nms.IpServer+":"+strconv.Itoa(int(cmdCheck.Nms.PortTl1)))
			msg := "Koneksi Ke NMS Timeout"
			isErr = service.doSendIfError(request, msg, err)
			if isErr {
				break
			}

			reply := make([]byte, 10240)
			if strings.ToUpper(cmdCheck.Nms.Vendor) != strings.ToUpper("ZTE") {
				service.doCommandTl1(cmdCheck.Login, conn, reply, &results, request)
			}
			for _, command := range cmdCheck.CmdList.Cmds {
				service.doCommandTl1(command, conn, reply, &results, request)
			}
			service.doCloseTl1(cmdCheck.Logout, conn, cmdCheck.Nms.Vendor, &results, request)
		}
		dataBot = strings.Join(results, "")
		service.DoSend(request, dataBot)
	}
}

func (service *ConfigViaTelegramServiceImpl) doCommandSsh(cmdCheck tl12.DtoCmdCheck, command string, results *[]string) {
	sshConfig := &ssh.ClientConfig{
		User: cmdCheck.Nms.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(cmdCheck.Nms.Password),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}
	addr := cmdCheck.Nms.IpServer + ":" + strconv.Itoa(int(cmdCheck.Nms.PortTl1))
	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, gpon := range cmdCheck.Gpons {
		slot := ""
		port := ""
		if gpon.SlotPort != "" {
			slot = strings.Split(gpon.SlotPort, "-")[0]
			port = strings.Split(gpon.SlotPort, "-")[1]
		}
		*results = append(*results, "\nIP GPON: "+gpon.IpGpon)
		session, err := conn.NewSession()
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()

		in, err := session.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		out, err := session.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		err = session.Start("telnet " + gpon.IpGpon)
		if err != nil {
			log.Fatal(err)
		}

		var output []byte
		go func(in io.WriteCloser, out io.Reader, output *[]byte) {
			var (
				line string
				r    = bufio.NewReader(out)
				errs error
			)
			for {
				b, err := r.ReadByte()
				errs = err
				if errs != nil {
					break
				}

				line += string(b)

				if strings.Contains(line, ">>User password:") && strings.HasSuffix(line, ":") {
					_, err := in.Write([]byte(gpon.Password + "\n"))
					errs = err
					break
				} else if strings.Contains(line, ">>User name:") && strings.HasSuffix(line, ":") {
					_, errs = in.Write([]byte(gpon.Username + "\n"))
					if errs != nil {
						break
					}
				}
			}

			if errs == nil {
				for {
					b, err := r.ReadByte()
					errs = err
					if errs != nil {
						break
					}

					line += string(b)

					if !strings.Contains(line, ">enable") && strings.HasSuffix(line, ">") {
						_, errs = in.Write([]byte(command))
						if errs != nil {
							break
						}
					}
					if (strings.Contains(line, "#display ont autofind all\r") ||
						strings.Contains(line, "#display ont info by-sn "+gpon.OnuId+" \r") ||
						strings.Contains(line, "#display service-port port 0/"+slot+"/"+port+" ont "+gpon.OnuId+"")) &&
						!strings.Contains(line, "#quit\r") {
						*output = append(*output, b)
					}
				}
			}
		}(in, out, &output)
		_ = session.Wait()
		*results = append(*results, strings.ReplaceAll(string(output), "\r", ""))
		*results = append(*results, "\n##################################################")
	}
}

func (service *ConfigViaTelegramServiceImpl) CheckUnReg(ctx context.Context, request request.SendMessageRequest, dataText string) {
	if !service.checkTelegramId(ctx, request) {
		return
	}
	ipAddresses := strings.Split(strings.TrimSpace(dataText), "\n")
	if len(ipAddresses) < 2 {
		msg := `Perintah Salah, Masukkan Data Dengan Benar

contoh:
/unreg
172.21.x.x
172.29.x.x
172.28.x.x
`
		err := errors.New(msg)
		_ = service.doSendIfError(request, msg, err)
		return
	}
	ipAddresses = helper.RemoveIndexSliceString(ipAddresses, 0)
	protocol, gpons, ok := service.checkGpon(ctx, request, ipAddresses)
	if !ok {
		return
	}

	ipAddresses = make([]string, 0)
	var configList []tl12.DtoConfig
	for index, gpon := range gpons {
		if index < 5 {
			var nms web.NmsRequest
			err := dto.Map(&nms, gpon.Nms)
			helper.PanicIfError(err)

			config := tl12.DtoConfig{
				IpGpon:         gpon.IpAddress,
				Username:       gpon.Username,
				Password:       gpon.Password,
				DefaultUserNms: gpon.DefaultUserNms,
				Nms:            nms,
			}

			ipAddresses = append(ipAddresses, "<b>"+gpon.Hostname+"("+gpon.IpAddress+")</b>")
			configList = append(configList, config)
		}
	}
	dataBot := "Tunggu Proses Cek Unreg Selesai, Maksimal Pengecekan 5 Perangkat GPON\n" + strings.Join(ipAddresses, "\n")
	service.DoSend(request, dataBot)
	go service.doProcessCheck(protocol, configList, dataBot, request, "UNREG")
}

func (service *ConfigViaTelegramServiceImpl) CheckReg(ctx context.Context, request request.SendMessageRequest, dataText string) {
	//TODO implement me
	panic("implement me")
}

func (service *ConfigViaTelegramServiceImpl) CheckService(ctx context.Context, request request.SendMessageRequest, dataText string) {
	//TODO implement me
	panic("implement me")
}

func (service *ConfigViaTelegramServiceImpl) ProcessService(ctx context.Context, request request.SendMessageRequest, dataText string) {
	//TODO implement me
	panic("implement me")
}

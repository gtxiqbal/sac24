package service

import (
	"context"
	"fmt"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/model/web/telegram/request"
	"strconv"
	"strings"
)

type TelegramServiceImpl struct {
	ConfigViaTelegramService
}

func NewTelegramServiceImpl(configViaTelegramService ConfigViaTelegramService) TelegramService {
	return &TelegramServiceImpl{ConfigViaTelegramService: configViaTelegramService}
}

func (service *TelegramServiceImpl) DoPostByTelegram(ctx context.Context, telegramRequest *request.TelegramRequest) web.Response[request.ResultResponse] {
	response := web.Response[request.ResultResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Mengirim data Telegram",
	}

	requestReply := request.SendMessageRequest{
		ChatId:           telegramRequest.Message.Chat.Id,
		ParseMode:        "HTML",
		ReplyToMessageId: telegramRequest.Message.MessageId,
	}
	text := telegramRequest.Message.Text
	if strings.TrimSpace(text) != "" {
		if strings.HasPrefix(text, "/") {
			dataBot := ""
			if strings.HasPrefix(text, "/start") {
				dataBot = fmt.Sprintf(`Selamat Datang <b>%s %s</b>
ID Telegram Anda: <code>%s</code>
Username: <b>@%s</b>
Gunakan /help untuk bantuan

Harap ID Telegram anda sudah didaftakan oleh Admin`,
					telegramRequest.Message.Chat.FirstName,
					telegramRequest.Message.Chat.LastName,
					strconv.Itoa(telegramRequest.Message.Chat.Id), telegramRequest.Message.Chat.Username)
				service.ConfigViaTelegramService.DoSend(requestReply, dataBot)
			} else if strings.HasPrefix(text, "/help") {
				dataBot = `Daftar Perintah

/unreg - <b>Cek Unregister ONT By Ip GPON</b>

contoh:
/unreg
172.21.x.x
172.29.x.x
172.28.x.x
`
				service.ConfigViaTelegramService.DoSend(requestReply, dataBot)
			} else if strings.HasPrefix(text, "/unreg") {
				service.ConfigViaTelegramService.CheckUnReg(ctx, requestReply, text)
			} else {
				dataBot = "Perintah Tidak Terdaftar"
				response.Message = dataBot
				service.ConfigViaTelegramService.DoSend(requestReply, dataBot)
			}
		}
	}

	return response
}

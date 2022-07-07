package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/model/web/telegram/request"
)

type TelegramService interface {
	DoPostByTelegram(ctx context.Context, telegramRequest *request.TelegramRequest) web.Response[request.ResultResponse]
}

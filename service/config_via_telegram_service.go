package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web/telegram/request"
)

type ConfigViaTelegramService interface {
	DoSend(requestReply request.SendMessageRequest, dataBot string)
	CheckUnReg(ctx context.Context, request request.SendMessageRequest, dataText string)
	CheckReg(ctx context.Context, request request.SendMessageRequest, dataText string)
	CheckService(ctx context.Context, request request.SendMessageRequest, dataText string)
	ProcessService(ctx context.Context, request request.SendMessageRequest, dataText string)
}

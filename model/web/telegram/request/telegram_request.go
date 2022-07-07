package request

type TelegramRequest struct {
	UpdateId int            `json:"update_id"`
	Message  MessageRequest `json:"message"`
}

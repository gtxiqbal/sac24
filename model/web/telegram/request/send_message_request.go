package request

type SendMessageRequest struct {
	ChatId           int    `json:"chat_id"`
	Document         string `json:"document"`
	Text             string `json:"text"`
	ParseMode        string `json:"parse_mode"`
	ReplyToMessageId int    `json:"reply_to_message_id"`
}

package request

type MessageRequest struct {
	MessageId      int                 `json:"message_id"`
	Form           FromRequest         `json:"form"`
	Chat           ChatRequest         `json:"chat"`
	Date           int                 `json:"date"`
	ReplyToMessage MessageRequestReply `json:"reply_to_message"`
	Text           string              `json:"text"`
}

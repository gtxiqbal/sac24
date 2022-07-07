package request

type MessageRequestReply struct {
	MessageId int         `json:"message_id"`
	Form      FromRequest `json:"form"`
	Chat      ChatRequest `json:"chat"`
	Date      int         `json:"date"`
	Text      string      `json:"text"`
}

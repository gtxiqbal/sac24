package request

type ResultResponse struct {
	Ok          bool           `json:"ok"`
	ErrorCode   int            `json:"error_code"`
	Description string         `json:"description"`
	Result      MessageRequest `json:"result"`
}

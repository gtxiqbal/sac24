package web

type Response[T any] struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

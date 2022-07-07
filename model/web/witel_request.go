package web

type WitelRequest struct {
	Id       string          `json:"id"`
	Alias    string          `json:"alias"`
	Nama     string          `json:"nama"`
	Regional RegionalRequest `json:"regional"`
}

package web

type StoRequest struct {
	Id    string       `json:"id"`
	Alias string       `json:"alias"`
	Nama  string       `json:"nama"`
	Witel WitelRequest `json:"witel"`
}

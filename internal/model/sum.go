package model

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Sum     int    `json:"sum"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

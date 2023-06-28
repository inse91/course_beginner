package model

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Sum int `json:"sum"`
	CommonResponse
}

type CommonResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

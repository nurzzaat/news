package models

type TokenResponse struct {
	Token interface{} `json:"token"`
}

type SuccessResponse struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Result string `json:"error"`
}

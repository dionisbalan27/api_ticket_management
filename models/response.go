package models

type MsgRes struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}

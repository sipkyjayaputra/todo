package dto

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}

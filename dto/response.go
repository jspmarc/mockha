package dto

type Response struct {
	Data         interface{} `json:"data"`
	ErrorMessage *string     `json:"error_message"`
	ServerTime   int64       `json:"server_time"`
}

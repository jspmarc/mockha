package dto

type Response struct {
	Data         interface{} `json:"data"`
	ErrorMessage *string     `json:"errorMessage"`
	ServerTime   int64       `json:"serverTime"`
}

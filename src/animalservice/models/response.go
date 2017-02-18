package models

type Response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

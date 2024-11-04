package models

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Error   any  `json:"error"`
}
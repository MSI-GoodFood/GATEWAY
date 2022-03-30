package models

type JSONResponseError struct {
	Code  string `json:"code" binding:"required"`
	Error string `json:"error"`
}

type JSONResponseSuccess struct {
	Code  string `json:"code" binding:"required"`
	Data  string `json:"data"`
}

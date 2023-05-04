package model

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type WebResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Errors []ErrorMsg  `json:"errors"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

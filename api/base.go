package api

type BaseResponse struct {
	Success int         `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data" yaml:"data"`
}

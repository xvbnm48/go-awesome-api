package dto

type ApiResponse[t any] struct {
	ResponseKey string `json:"response_key"`
	ResponseMsg string `json:"response_msg"`
	Data        T      `json:"data"`
}

package client

import "encoding/json"

// 客户端请求
type ClientRequest struct {
	Url  string      `json:"url"`
	Data interface{} `json:"Data"`
}

// 客户端响应
type ClientResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewClientResponse(code int, msg string, data interface{}) *ClientResponse {
	return &ClientResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewOkClientRes(data interface{}) *ClientResponse {
	return &ClientResponse{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}
}

func NewErrClientRes(msg string, data interface{}) *ClientResponse {
	return &ClientResponse{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}

func (r *ClientResponse) GetByte() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ClientResponse) GetString() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

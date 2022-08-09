package model

import (
	"encoding/json"
)

const successMsg = "成功"
const failedMsg = "失败"
const successStatus = "successful"
const failedStatus = "failed"
const successCode = 200
const failedCode = 500

type BaseModel struct {
}

type ApiResponse struct {
	BaseModel
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
}

func (response *ApiResponse) Successful(data interface{}) *ApiResponse {
	response.Msg = successMsg
	response.Status = successStatus
	response.Code = successCode
	response.Data = data
	return response
}

func (response *ApiResponse) Failed(data interface{}) *ApiResponse {
	response.Msg = failedMsg
	response.Status = failedStatus
	response.Code = failedCode
	response.Data = data
	return response
}

func (response *ApiResponse) Message(msg string) *ApiResponse {
	response.Msg = msg
	return response
}

func (response *ApiResponse) ToString() (string, error) {
	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), err
}

type DeviceConfig struct {
	BaseModel
	DeviceId   string           `json:"deviceId"`
	DeviceName string           `json:"deviceName"`
	Commands   []CommandContent `json:"command"`
}

type CommandContent struct {
	BaseModel
	Description string `json:"description"`
	Content     string `json:"content"`
}

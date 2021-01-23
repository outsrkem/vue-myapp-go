package models

import "time"

type GeneralErrorStruct struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

type resMessage struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

// 无数据返回体的统一消息
func NewResMessage(code string, msg string) *resMessage {
	return &resMessage{
		Code: code,
		Msg:  msg,
		//RequestTime: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		RequestTime: time.Now().UnixNano(),
	}
}

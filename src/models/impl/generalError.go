package impl

type GeneralErrorStruct struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	RequestTime int64 `json:"requestTime"`
}
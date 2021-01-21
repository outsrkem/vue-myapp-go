package models

// 注册用户的body结构
type UserRegisterInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册返回体结构
type UserRegisterStruct struct {
	MetaInfo metaInfo `json:"metaInfo"`
	Response response `json:"response"`
}
type metaInfo struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	RequestTime int64 `json:"requestTime"`
}
type response struct {
	Userid string `json:"userid"`
	Username string `json:"username"`
}
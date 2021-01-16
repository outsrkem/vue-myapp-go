package impl
type UserLoginStruct struct {
	MetaInfo MetaInfo `json:"metaInfo"`
	Response Response `json:"response"`
}
type MetaInfo struct {
	Status string `json:"status"`
	Msg string `json:"msg"`
	RequestTime int64 `json:"requestTime"`
}
type Response struct {
	Userid int `json:"userid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role int `json:"role"`
	Expires int8 `json:"expires"`
	Token string `json:"token"`
}

// 登录时提交的数据
type UserLoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
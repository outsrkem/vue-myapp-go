package impl

type UserLoginStruct struct {
	MetaInfo MetaInfo `json:"metaInfo"`
	Response Response `json:"response"`
}
type MetaInfo struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}
type Response struct {
	Userid   string    `json:"userid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string    `json:"role"`
	Expires  string   `json:"expires"`
	Token    string `json:"token"`
}

// 登录时提交的数据
type UserLoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

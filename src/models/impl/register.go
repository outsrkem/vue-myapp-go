package impl

// 注册用户的body结构
type UserRegisterInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
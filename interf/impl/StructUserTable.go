package impl

type UserTable struct {
	Id         int    `json:"id"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	CreatTime  int64  `json:"creat_time"`
	UpdateTime int64  `json:"update_time"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	Token      string `json:"token"`
}

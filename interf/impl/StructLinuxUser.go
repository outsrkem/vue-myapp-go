package impl

/*
	定义Linux用户存入数据库的结构体
*/
type LinuxListUser struct {
	User string `json:"user"`
	Ip   string `json:"ip"`
	Pwd  string `json:"pwd"`
	Port string `json:"port"`
}

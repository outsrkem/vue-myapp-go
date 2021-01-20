package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
	"menu/middle"
)

func (u *UserTable) UserLogin() string {
	data := db.Get([]byte(u.UserName), db.UserTable)
	if data == nil {
		fmt.Println(u.UserName + "账号不存在")
		return u.UserName + "账号不存在"
	}

	var user UserTable
	err := json.Unmarshal(data.Value, &user)
	if err != nil {
		fmt.Println(err)
		return "反序列化失败"
	}

	if u.Password == user.Password {
		u.Password = ""
		u.UserStatus = user.UserStatus
		token := middle.Token(user.UserName, user.Role)
		u.Token = token
	} else {
		return "u.UserName + 密码错误"
	}

	return "成功"
}

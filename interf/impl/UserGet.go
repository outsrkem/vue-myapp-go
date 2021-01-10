package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
)

func (u *UserTable) UserGet(username string) {

	//获取数据库数据
	data := db.Get([]byte(username), db.UserTable)
	if data == nil {
		u.Status = "UserGet获取数据库错误"
		return
	}

	//反序列化数据
	var user UserTable
	err := json.Unmarshal(data.Value, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	u.UserName = user.UserName
	u.Password = user.Password
	u.Role = user.Role
	u.Status = user.Status
	u.CreatTime = user.CreatTime
	u.UpdateTime = user.UpdateTime
}

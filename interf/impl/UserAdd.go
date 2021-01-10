package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
	"time"
)

func (u *UserTable) UserAdd() {

	//判断是否是第一次添加数据，如果存在，则只修改更新时间字段
	data := db.Get([]byte(u.UserName), db.UserTable)
	if data != nil {
		var user UserTable
		err := json.Unmarshal(data.Value, &user)
		if err != nil {
			fmt.Println(err)
			return
		}
		u.CreatTime = user.CreatTime
		u.UpdateTime = time.Now().Unix()
	} else {
		u.CreatTime = time.Now().Unix()
		u.UpdateTime = time.Now().Unix()
	}

	//将UserTable对象序列化为json格式[]byte数据当作数据库value
	val, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		return
	}

	//用户名作为key添加用户数据
	db.Add([]byte(u.UserName), val, db.UserTable)
}

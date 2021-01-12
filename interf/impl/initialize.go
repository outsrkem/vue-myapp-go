package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
	"time"
)

/*
	初始化账号密码
*/
func Initialize(adminUser, adminPassword string) {

	admin := UserTable{
		UserName:   adminUser,
		Password:   adminPassword,
		Status:     "0",
		Role:       "admin",
		CreatTime:  time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	user, err := json.Marshal(&admin)
	if err != nil {
		fmt.Println(err)
	}

	//如果数据库为空或账号有变动，则执行数据库修改
	data := db.Get([]byte(admin.UserName), db.UserTable)
	if data == nil || admin.UserName != "admin" || admin.Password != "admin" {
		db.Add([]byte(admin.UserName), user, db.UserTable)
	}
}

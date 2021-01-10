package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
	"time"
)

/*
	根据指定参数，添加数据
*/

func (l *LinuxList) Add(user *LinuxListUser) {
	//查看是否能连接Linux服务器，返回err不能连接
	_, err := GetLinuxCmd(user)
	if err != nil {
		fmt.Println(err)
		l.MetaInfo.Msg = "无法连接linux服务器"
		l.MetaInfo.Status = 502
		l.MetaInfo.RequestTime = time.Now().Unix()
		return
	}

	//序列化成json格式[]byte类型
	val, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	l.MetaInfo.Msg = "success"
	l.MetaInfo.Status = 201
	l.MetaInfo.RequestTime = time.Now().Unix()

	key := []byte(user.Ip)
	//存入数据库
	db.Add(key, val, db.LinuxList)
}

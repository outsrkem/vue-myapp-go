package impl

import (
	"menu/db"
	"time"
)

/*
	根据指定key删除数据库对应数据
*/
func (l *LinuxList) Del(ip string) {
	//删除数据库数据
	db.Del([]byte(ip), db.LinuxList)

	l.MetaInfo.Status = 200
	l.MetaInfo.Msg = "success"
	l.MetaInfo.RequestTime = time.Now().Unix()

}

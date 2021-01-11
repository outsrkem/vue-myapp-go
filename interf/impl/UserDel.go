package impl

import "menu/db"

func (u *UserTable) UserDel(username string) {
	db.Del([]byte(username), db.UserTable)
	u.UserName = username
	u.Status = "删除成功"
}

package interf

import "menu/interf/impl"

/*
	操作登录用户接口
*/
type UserTableI interface {
	UserAdd()
	UserDel(username string)
	UserGet(username string)
	UserGetAll() *[]impl.UserTable
	UserLogin() string
}

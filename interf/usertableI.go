package interf

import "menu/interf/impl"

/*
	操作登录用户接口
*/
type UserTableI interface {
	UserAdd()
	UserDel(username string)
	UserGetAll() *[]impl.UserTable
	UserLogin() string
	UserGet(username string) *[]impl.UserTable
	UserPageGet(pagesize, pagenum string) (*[]impl.UserTable, int, int)
	//UserPageNumGet(pagesize,pagenum string) *[]impl.UserTable
}

package interf

import "menu/interf/impl"

type UserTableI interface {
	UserAdd()
	UserDel(username string)
	UserGet(username string)
	UserGetAll() *[]impl.UserTable
	UserLogin() string
}

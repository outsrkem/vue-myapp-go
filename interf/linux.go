package interf

import (
	"menu/interf/impl"
)

/*
	定义操作Linux的相关接口
*/
type LinuxInterface interface {
	Add(user *impl.LinuxListUser)
	GetAll()
	Del(ip string)
}

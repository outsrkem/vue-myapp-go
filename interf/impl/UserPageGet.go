package impl

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"menu/db"
)

/*
	返回分页信息
*/
func (u *UserTable) UserPageGet(pagesize, pagenum string) (*[]UserTable, int, int) {
	//查询数据库
	index := db.GetIndex(db.UserTable)
	data := db.GetAll(db.UserTable)
	if data == nil {
		fmt.Println("数据库查询为空")
		return nil, 0, 0
	}

	//序列化数据库数据
	dataList := make([]UserTable, index)
	for i, k := range data {
		err := json.Unmarshal(k.Value, &u)
		if err != nil {
			fmt.Println("序列化失败")
			return nil, 0, 0
		}
		u.Id = i + 1
		dataList[i] = *u
	}

	//size为分页页数，num为当前第几页
	size := cast.ToInt(pagesize)
	num := cast.ToInt(pagenum)

	//如果num的值为1，表示只修改分页，返回分页数和数据总数，如果num不等于1，跳转到第几页，返回该页数据
	if num == 1 {
		if size >= index {
			return &dataList, 1, index
		} else if size < index {
			num := (index / size) + 1
			userlist := dataList[0:size]
			return &userlist, num, index
		} else {
			return nil, 0, 0
		}
	} else {
		if size*num >= index {
			userlist := dataList[size*num-size:]
			return &userlist, 0, 0
		} else if size*num < index {
			userlist := dataList[size*num-size : size*num]
			return &userlist, 0, 0
		} else {
			return nil, 0, 0
		}
	}
}

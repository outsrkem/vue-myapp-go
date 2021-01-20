package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
)

/*
	查询全部用户信息
*/
func (u *UserTable) UserGetAll() *[]UserTable {

	index := db.GetIndex(db.UserTable)
	data := db.GetAll(db.UserTable)
	if data == nil {
		fmt.Println("UserGetAll数据库无数据")
		return nil
	}

	usertable := make([]UserTable, index)
	var userT UserTable
	for i, k := range data {
		err := json.Unmarshal(k.Value, &userT)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		userT.Id = i + 1
		usertable[i] = userT
	}

	return &usertable
}

package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
)

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
		usertable[i].Id = i + 1
		usertable[i].UserName = userT.UserName
		usertable[i].Password = userT.Password
		usertable[i].Status = userT.Status
		usertable[i].Role = userT.Role
		usertable[i].CreatTime = userT.CreatTime
		usertable[i].UpdateTime = userT.UpdateTime
	}

	return &usertable
}

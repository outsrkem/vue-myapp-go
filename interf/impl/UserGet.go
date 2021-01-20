package impl

import (
	"encoding/json"
	"fmt"
	"menu/db"
)

/*
	前缀查询逻辑
*/
func (u *UserTable) UserGet(username string) *[]UserTable {
	data := db.PrefixGet(db.UserTable, username)
	if data == nil {
		fmt.Println("前缀查询错误")
		return nil
	}

	userList := make([]UserTable, 0, 5)
	for i, entry := range data {
		err := json.Unmarshal(entry.Value, &u)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		u.Id = i + 1
		userList = append(userList, *u)
	}
	return &userList
}

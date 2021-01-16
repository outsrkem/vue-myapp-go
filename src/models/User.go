package models

import (
	"fmt"
	"mana/src/connections/database/mysql"
	"strconv"
	"time"
)

// 用户表
type UserInfo struct {
	ID         string `json:"id"`         // id
	USERID     string `json:"name"`       // 用户id
	USERNAME   string `json:"username"`   // 用户名
	NICKNAME   string `json:"nickname"`   // 昵称
	ROLE       string `json:"role"`       // 角色
	PASSWD     string `json:"passwd"`     // 密码
	UPDATETIME string `json:"change"`     // 最近一次密码修改时间
	EXPIRES    string `json:"expires"`    // 密码过期时间
	INACTIVE   string `json:"inactive"`   // 用户状态
	CREATETIME string `json:"createtime"` // 创建时间
}

func InstUser(name string, passwd string) string {
	atTimes := time.Now().Unix()
	atTimesStr := time.Unix(atTimes, 0).Format("2006-01-02 15:04:05")
	uid := time.Now().UnixNano()
	nickname := name
	role, expires, inactive := 3, 2, 1
	sqlStr := `INSERT INTO user (USERID, USERNAME, NICKNAME, ROLE, PASSWD, UPDATETIME, EXPIRES, INACTIVE, CREATETIME) VALUES (?,?,?,?,?,?,?,?,?);`
	_, err := mysql.DB.Exec(sqlStr, uid, name, nickname, role, passwd, atTimesStr, expires, inactive, atTimesStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return "123"
	}
	// 转换为string类型
	return strconv.FormatInt(uid, 10)
}

// 查询单条
func SelectUserQueryRow(username string) (*UserInfo, error) {
	var u UserInfo
	sqlStr := `SELECT ID,USERID,USERNAME,NICKNAME,ROLE,PASSWD,EXPIRES,INACTIVE,CREATETIME,UPDATETIME FROM  user WHERE USERNAME = ?`
	var row = mysql.DB.QueryRow(sqlStr, username)
	//err := row.Scan(u.ID, u.USERID, u.USERNAME, u.NICKNAME, u.ROLE, u.PASSWD, u.UPDATETIME, u.EXPIRES, u.INACTIVE, u.CREATETIME)
	err := row.Scan(&u.ID, &u.USERID, &u.USERNAME, &u.NICKNAME, &u.ROLE, &u.PASSWD, &u.EXPIRES, &u.INACTIVE, &u.CREATETIME, &u.UPDATETIME)
	if err != nil {
		fmt.Println("asd", err.Error())

	}
	return &u, err
}

// 查询多条
func SelectUserQueryMultiRow(id int) {
	// 1.sql
	sqlStr := `SELECT * FROM  user WHERE ID > ?`
	// 2.执行
	rows, err := mysql.DB.Query(sqlStr, id)
	if err != nil {
		fmt.Println("exec %s query failed, err:%v\n", sqlStr, err)
	}
	// 3 一定要关闭连接
	defer rows.Close()
	// 4. 循环取值
	for rows.Next() {
		var u UserInfo
		err := rows.Scan(&u.ID, &u.USERID, &u.USERNAME, &u.NICKNAME, &u.ROLE, &u.PASSWD, &u.UPDATETIME, &u.EXPIRES, &u.INACTIVE, &u.CREATETIME)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Printf("u:%#v\n", u)
	}

}

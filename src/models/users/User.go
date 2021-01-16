package users

import (
	"fmt"
	"mana/src/connections/database/mysql"
	"strconv"
	"time"
)

// 用户表
type userInfo struct {
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

type userCenter struct {
	ID         string `json:"id"`         // id
	USERID     string `json:"name"`       // 用户id,外键
	USERNAME   string `json:"username"`   // 用户名，外键
	NICKNAME   string `json:"nickname"`   // 昵称
	MOBILE     string `json:"mobile"`     //手机
	EMAIL      string `json:"email"`      // 邮箱
	DESCRIBES  string `json:"describes"`  // 描述说明
	PICTURE    string `json:"picture"`    // 头像
	CREATETIME string `json:"createtime"` // 创建时间
	UPDATETIME string `json:"updatetime"` // 最近更新时间
}

// 注册用户
func InstUser(name string, passwd string) (string, error) {
	atTimes := time.Now().Unix()
	atTimesStr := time.Unix(atTimes, 0).Format("2006-01-02 15:04:05")
	uid := time.Now().UnixNano()
	nickname := name
	role, expires, inactive := 3, 2, 1
	tx, err := mysql.DB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("事务开启失败:%v\n", err)
		return "", err
	}
	// 插入用户表信息
	sqlStr1 := `INSERT INTO user (USERID, USERNAME, NICKNAME, ROLE, PASSWD, UPDATETIME, EXPIRES, INACTIVE, CREATETIME) VALUES (?,?,?,?,?,?,?,?,?);`
	_, err = tx.Exec(sqlStr1, uid, name, nickname, role, passwd, atTimesStr, expires, inactive, atTimesStr)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("用户表插入失败:%v\n", err)
		return "", err
	}
	// 插入用户中心表信息
	sqlStr2 := `INSERT INTO user_center (USERID, USERNAME, NICKNAME, CREATETIME, UPDATETIME) VALUES (?,?,?,?,?);`
	_, err = tx.Exec(sqlStr2, uid, name, nickname, atTimesStr, atTimesStr)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("用户中心表插入失败:%v\n", err)
		return "", err
	}
	// 提交事务
	if err = tx.Commit(); err != nil {
		// 事务回滚
		tx.Rollback()
		fmt.Println("事务回滚...")
		return "", err
	}

	// 转换为string类型
	return strconv.FormatInt(uid, 10), err
}

// 查询单条
func SelectUserQueryRow(username string) (*userInfo, error) {
	var u userInfo
	sqlStr := `SELECT ID,USERID,USERNAME,NICKNAME,ROLE,PASSWD,EXPIRES,INACTIVE,CREATETIME,UPDATETIME FROM  user WHERE USERNAME = ?`
	fmt.Println(sqlStr)
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
		var u userInfo
		err := rows.Scan(&u.ID, &u.USERID, &u.USERNAME, &u.NICKNAME, &u.ROLE, &u.PASSWD, &u.UPDATETIME, &u.EXPIRES, &u.INACTIVE, &u.CREATETIME)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Printf("u:%#v\n", u)
	}

}

// 查询单条
func SelectUidUserQueryRow(uid string) (*userInfo, error) {
	var u userInfo
	sqlStr := `SELECT ID,USERID,USERNAME,NICKNAME,ROLE,PASSWD,EXPIRES,INACTIVE,CREATETIME,UPDATETIME FROM  user WHERE USERID = ?`
	var row = mysql.DB.QueryRow(sqlStr, uid)
	err := row.Scan(&u.ID, &u.USERID, &u.USERNAME, &u.NICKNAME, &u.ROLE, &u.PASSWD, &u.EXPIRES, &u.INACTIVE, &u.CREATETIME, &u.UPDATETIME)
	if err != nil {
		fmt.Println("asd", err.Error())

	}
	return &u, err
}

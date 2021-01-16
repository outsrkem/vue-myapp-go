package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var USERNAME = "abc"
var PASSWORD = "123456"
var NETWORK = "tcp"
var SERVER = "10.10.10.10"
var PORT = 3306
var DATABASE = "mana"
var MaxOpenConn = 10
var MaxIdleConn = 5

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=true", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("a连接失败", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("b连接失败", err)
		return
	}
	// 配置连接池最大连接数
	DB.SetMaxOpenConns(MaxOpenConn)
	// 配置连接池最大空闲连接数
	DB.SetMaxIdleConns(MaxIdleConn)
	return
}

package mysql

import (
	"database/sql"
	"fmt"
	"mana/src/config"

	// mysql 连接驱动
	_ "github.com/go-sql-driver/mysql"
)

// DB mysql
var DB *sql.DB
var _log = config.Log()

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) (err error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=true",
		cfg.DbUser, cfg.DbPassword, cfg.DbNetwork, cfg.DbServer, cfg.DbPort, cfg.DbName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		_log.Error("数据库配置错误,", err)
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		_log.Error("数据库连接失败,", err)
		panic(err)
	}
	// 配置连接池最大连接数
	DB.SetMaxOpenConns(cfg.DbMaxOpenConn)
	// 配置连接池最大空闲连接数
	DB.SetMaxIdleConns(cfg.DbMaxIdleConn)
	// 连接成功后清空密码
	cfg.DbPassword = ""
	return
}

package config

import (
	"flag"
)

// Config 配置结构体
type Config struct {
	Listen        string `json:"listen"`
	Port          string `json:"port"`
	DbUser        string `json:"db_user"`
	DbPassword    string `json:"db_password"`
	DbNetwork     string `json:"db_network"`
	DbServer      string `json:"db_server"`
	DbPort        string `json:"db_port"`
	DbName        string `json:"db_name"`
	DbMaxIdleConn int    `json:"db_MaxIdleConn"`
	DbMaxOpenConn int    `json:"db_MaxOpenConn"`
}

// InitConfig 初始化配置
func InitConfig() *Config {
	var _cfg Config
	flag.StringVar(&_cfg.Listen, "bind-address", "0.0.0.0", "bind-address")
	flag.StringVar(&_cfg.Port, "secure-port", "9443", "secure-port")
	flag.StringVar(&_cfg.DbUser, "db-username", "abc", "db-username")
	flag.StringVar(&_cfg.DbPassword, "db-password", "123456", "db-password")
	flag.StringVar(&_cfg.DbNetwork, "db-network", "tcp", "db-network")
	flag.StringVar(&_cfg.DbServer, "db-server", "127.0.0.1", "db-server")
	flag.StringVar(&_cfg.DbPort, "db-port", "3306", "db-port")
	flag.StringVar(&_cfg.DbName, "db-name", "mana", "db-name")
	flag.IntVar(&_cfg.DbMaxIdleConn, "db-maxidleconn", 5, "DbMaxIdleConn")
	flag.IntVar(&_cfg.DbMaxOpenConn, "db-maxopenconn", 10, "DbMaxOpenConn")

	// 解析命令行参数
	flag.Parse()
	return &_cfg
}

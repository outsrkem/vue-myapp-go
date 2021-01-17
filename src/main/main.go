package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mana/src/config"
	"mana/src/connections/database/mysql"
	"mana/src/routers"
)

func main() {
	// 加载日志
	log := config.Log()

	// 加载配置文件
	//cfg, err := config.ParseConfig("../config/config.json")
	//if err != nil {
	//	log.Error("配置文件读取错误")
	//	panic(err.Error())
	//}

	// 加载默认配置
	cfg := config.ParseConfig()

	// 连接数据库MySql
	mysql.InitDB(cfg)

	// 关掉控制台颜色
	gin.DisableConsoleColor()
	// gin.Default 通过New创建了Engine实例， 并Use了 Logger Recovery两个HandlerFunc中间件。
	r := gin.Default()
	// 服务路由请求
	routers.Index(r)

	// 服务运行
	err1 := r.Run(cfg.Listen + ":" + cfg.Port)
	if err1 != nil {
		log.Error("服务运行错误")
	}
}

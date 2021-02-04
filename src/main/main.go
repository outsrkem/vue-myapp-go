package main

import (
	"mana/src/config"
	"mana/src/connections/database/mysql"
	"mana/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载日志
	log := config.Log()

	// 加载默认配置
	cfg := config.InitConfig()

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

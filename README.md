# vue-myapp-go

```shell
.
├── Makefile
├── README.md
├── cli                     
│   └── cli.go
├── config                      全局配置
│   ├── connections.go
│   ├── cookie.go
│   ├── jwt.go
│   └── env.go
├── connections                 存储连接
│   ├── database
│   │   ├── mongodb
│   │   └── mysql
│   └── redis
│       └── redis.go
├── controllers                 控制器
│   └── MainController.go
├── filters                     中间件
│   ├── auth                    认证中间件
│   │   ├── drivers             认证引擎
│   │   └── auth.go   
│   └── filter.go              
├── frontend                    前端资源
│   ├── assets
│   │   ├── css
│   │   ├── images
│   │   └── js
│   ├── dist
│   └── templates
│       └── index.tpl
├── handle.go                   全局错误处理
├── main.go                     
├── models                      模型
│   └── User.go
├── module                      项目模块
│   │── schedule
│   │   └── schedule.go         定时任务模块
│   │── logger
│   │   └── logger.go 
│   └── server
│       └── server.go           
├── routers                     路由
│   └── api_routers.go          
├── routers.go                  路由初始化设置
├── routers_test.go             api测试
├── storage                     
│   ├── cache                   缓存文件
│   └── logs                    项目日志
│       ├── access.log 
│       ├── info.log          
│       └── error.log
└── vendor                      govendor 第三方包
```


# vue-myapp-go

## **介绍**

本项目为学习golang所开发的自用实验项目，功能比较简单，主要是获取Linux服务器性能和k8s集群的信息。

## 项目介绍

### 使用

下载

```bash
git clone https://gitee.com/Outsrkem/vue-myapp-go.git
```

安装

windows：

```go
//设置windows编译环境
SET CGO_ENABLED=0	// 禁用CGO
SET GOOS=windows  	// 目标平台是linux
SET GOARCH=amd64  	// 目标处理器架构是amd64

//编译
go build -o app.exe

//运行
app.exe
```

linux：

```go
//设置linux编译环境
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

//编译
go build -o app

//运行
chmod u+x app
./app
```

### 目录说明

- .idea：开发工具为golang所生成的目录
- db：操作nutsdb数据库的增删改查方法
- interf：接口目录
  - impl：为实现接口的方法
- middle：中间件目录
- router：路由目录
  - method：路由调用的具体方法目录
- go.mod：mod依赖管理，包含所有使用到的模块
- go.sum：记录所依赖的项目的版本的锁定
- main.go：程序主函数
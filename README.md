# vue-myapp-go

## **介绍**

本项目为学习golang所开发的自用实验项目，功能比较简单，主要是获取、添加、删除Linux服务器性能和k8s集群的信息，以及登录校验token等功能，使用go语言版本为go 1.15。

## 项目介绍

### 使用

**下载**

```bash
git clone https://gitee.com/Outsrkem/vue-myapp-go.git
```

**安装**

需要go语言环境

windows：

```bash
# 设置windows编译环境
SET CGO_ENABLED=0	# 禁用CGO
SET GOOS=windows  	# 目标平台是linux
SET GOARCH=amd64  	# 目标处理器架构是amd64

# 编译
go build -o app.exe

# 运行
app.exe

# 查看运行参数
app.exe -h

Usage of menu.exe:
  -port int
        set server port (default 8080)			//设置端口
  -pwd string
        set admin password (default "admin")	//设置或修改管理员账号，如果没有指定，默认admin
  -user string
        set admin username (default "admin")	//设置或修改管理员密码，如果没有指定，默认admin
```

linux：

```bash
# 设置linux编译环境
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# 编译
go build -o app

# 运行
chmod u+x app
./app

# 查看运行参数
./app -h

Usage of menu.exe:
  -port int
        set server port (default 8080)			//设置端口
  -pwd string
        set admin password (default "admin")	//设置或修改管理员账号，如果没有指定，默认admin
  -user string
        set admin username (default "admin")	//设置或修改管理员密码，如果没有指定，默认admin

```

### 目录说明

- .idea：开发工具为golang所生成的目录
- db：操作nutsdb数据库的增删改查方法
- interf：接口目录
  - impl：为实现接口的方法
- middle：中间件目录
- router：路由目录
  - method：路由调用的具体方法目录
  - index.go：所有路由地址
- go.mod：mod依赖管理，包含所有使用到的模块
- go.sum：记录所依赖的项目的版本的锁定
- main.go：程序主函数

### 请求说明

**操作登录用户**

- 登录：/login?username=admin&password=admin
  - GET请求
  - 返回值token：其它所有操作均需要在header上添加`Token`（首字母大写）字段和Token值



- 查询所有登录用户：/api/v1/common/user/table
  - GET请求：需要`admin`或`general`权限



- 添加用户：/api/v1/common/user/table?username=huang&password=123123&role=general&status=0
  - POST请求：需要`admin`权限
  - 请求参数：role用户权限，分为admin和general，status为用户状态，暂时没有



- 删除用户：/api/v1/common/user/table?username=huang
  - DELETE请求：需要`admin`权限



**操作linux服务器信息**

- 添加linux账号：/api/v1/common/resource/monitor?user=root&ip=192.168.40.102&pwd=123123&port=22
  - POST请求：如果账号无法实际连接服务器，添加失败，需要`admin`权限



- 查看linuxl所有服务器资源：/api/v1/common/resource/monitor
  - GET请求：需要`admin`或`general`权限



- 删除linux服务器账号：/api/v1/common/resource/monitor?ip=192.168.40.102
  - DELETE请求：需要`admin`权限



**操作k8s集群信息**

- 添加k8s集群：/api/v1/common/kubernetes/cluster
  - POST请求：需要`admin`权限
  - 请求参数：body中的raw格式，添加json格式的k8sconfig集群配置文件



- 查看k8s所有集群：/api/v1/common/kubernetes/cluster?type=cluster
  - GET请求：需要`admin`或`general`权限
  - 请求参数：type为查看的类型，有cluster，namespaces，workingload



- 查看指定集群的命名空间（namespaces）：/api/v1/common/kubernetes/cluster?type=namespaces&address=集群连接地址
  - GET请求：`address`是k8sconfig配置文件中的集群连接地址



- 查看指定集群的deployment信息：/api/v1/common/kubernetes/cluster?type=workingload&address=集群连接地址&namespaces=default&control=deployments
  - GET请求：需要`admin`或`general`权限
  - 请求参数：其中namespaces值是名称空间名字，control和type值是固定的



- 删除k8s集群：/api/v1/common/kubernetes/cluster?address=集群连接地址
  - DELETE请求：需要`admin`权限


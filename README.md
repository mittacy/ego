# ego
快速创建 gin 项目的命令行工具

### 1. 安装

确保使用 go module

```shell
$ GO111MODULE=on
$ GOPROXY=https://goproxy.cn/,direct
```

安装

```shell
$ go get -u github.com/mittacy/ego@latest
```

### 2. 创建项目

通过 ego 命令创建项目模板：

```bash
$ ego new helloworld
```

输出:

```shell
├── bootstrap               # 初始化顺序调用封装
│   └── init.go
├── apierr                  # 服务错误码和错误定义
│   ├── code.go
│   └── err.go
├── pkg                     # 各种工具封装
│   ├── checker             # 校验器封装
│   │   └── validator.go
│   ├── config              # 配置全局初始化、配置结构
│   │   ├── init.go
│   │   ├── jwt.go
│   │   ├── log.go
│   │   ├── mysql.go
│   │   ├── redis.go
│   │   └── server.go
│   ├── jwt
│   │   └── token.go
│   ├── logger              # 日志封装
│   │   ├── init.go
│   │   └── utils.go
│   ├── response            # 响应封装
│   │   ├── log.go
│   │   └── response.go
│   └── store               # 持久化、缓存封装
│       ├── cache
│       │   ├── custom.go
│       │   └── redigo.go
│       └── db
│           └── gorm.go
├── app
│   ├── validator           # 数据请求、响应结构体定义以及参数校验
│   │   └── userValidator
│   │       └── user.go
│   ├── transform           # 响应数据处理、封装
│   │   └── userTransform
|   |        └── user.go
│   ├── api                 # api控制器，这里只进行请求解析、service编排与调用
│   │   └── user.go
│   ├── service             # 服务层，处理逻辑，实现api中各个服务接口
│   │   └── user.go
│   ├── data                # 数据存储层，实现service中各个data接口
│   │   └── user.go
│   └── model               # 定义与数据库的映射结构体
│       └── user.go
├── middleware              # 中间件
│   └── core.go
├── router                  # 路由初始化
│   ├── router.go
│   ├── wire.go             # 依赖注入，生成各种api控制器创建函数
│   └── wire_gen.go
└── utils                   # 工具封装
│   └── err.go
├── logs                    # 日志文件目录
│   ├── err
│   │   └── default.log
│   └── info
│       └── default.log
├── default.yaml
├── go.mod
├── go.sum
├── main.go
```

### 3. 创建业务接口

**需要在项目根目录运行**

```shell
$ ego tpl api article
```

将会创建 api、validator、transform、service、data、model 代码架构


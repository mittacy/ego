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
├── bootstrap					# 初始化顺序调用封装
│   └── init.go
├── apierr						# 服务错误码和错误定义
│   ├── code.go
│   └── err.go
├── pkg							# 各种外部工具封装
│   ├── cache					# 缓存封装
│   │   ├── config.go
│   │   └── redis.go
│   ├── config					# 配置全局初始化、配置结构
│   │   └── viper.go
│   ├── log						# 日志封装
│   │   ├── biz.go
│   │   ├── config.go
│   │   ├── log.go
│   │   ├── zap.go
│   │   └── zap_test.go
│   ├── mysql					# mysql封装
│   │   ├── config.go
│   │   └── gorm.go
│   └── response				# 响应封装
│       ├── log.go
│       ├── response.go
│       └── validator.go
├── cmd							# 服务
│   ├── api						# API监听服务
│   │   └── main.go
│   └── job						# 流式任务处理服务
│       └── main.go
├── interface				    # 对外服务
│   ├── api						# API控制器
│   │   └── user.go
│   ├── job						# 流式任务
│   │   └── exampleJob
│   │       ├── processor.go	# 任务处理器
│   │       └── task.go			# 生成任务
│   └── task					# 定时任务
├── internal					# 内部服务
│   └── validator				# 数据请求、响应结构体定义以及参数校验
│   │   └── userValidator
│   │       └── user.go
│   ├── transform				# 响应数据处理、封装
│   │   └── user.go
│   ├── service					# 服务层，处理逻辑
│   │   └── user.go
│   ├── data					# 数据查询、存储层
│   │   └── user.go
│   └── model					# 定义与数据库的映射结构体
│       └── user.go
├── middleware              	# 中间件
└── router						# 路由
│   ├── admin.go
│   ├── request.go
│   └── router.go
├── Makefile
├── main.go
```

### 3. 生成业务模板

**tpl命令需要在项目根目录运行**

```shell
# 创建 api、validator、transform、service、data、model 代码模板
$ ego tpl api article

# 创建 service、data、model 代码模板
$ ego tpl service article

# 创建 data、model 代码模板
$ ego tpl data article

# 创建定时任务 task 代码模板
$ ego tpl task notice

# 创建异步任务 job 代码模板
$ ego tpl job sendEmail
```


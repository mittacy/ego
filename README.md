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
├── bootstrap					# 服务依赖初始化
│   ├── http.go					# http服务
│   ├── job.go					# 异步任务服务
│   ├── log.go
│   ├── task.go					# 定时任务服务
│   └── viper.go
├── apierr						# 业务码定义
│   └── code.go
├── cmd
│   └── start					
│       ├── http				# http
│       │   └── http.go
│       ├── job					# 异步任务
│       │   └── job.go
│       └── task				# 定时任务
│       │   └── task.go
│		└── start.go			# 服务启动程序
├── bin							# 可执行文件存储目录，不上传
├── config						# 配置设置
│   ├── async.go
│   ├── async_config
│   │   └── job.go
│   ├── mysql.go
│   ├── redis.go
│   └── task.go
├── middleware					# 中间件
│   ├── requestLog.go
│   └── requestTrace.go
├── router						# 路由
│   ├── admin.go
│   └── router.go
├── app
│   ├── api						# api控制器
│   │   └── user.go
│   ├── internal
│   │   ├── data
│   │   │   └── user.go
│   │   ├── model
│   │   │   └── user.go
│   │   ├── service
│   │   │   └── user.go
│   │   ├── transform			# 响应数据处理、封装
│   │   │   └── user.go
│   │   └── validator			# 数据请求与参数校验、响应结构体
│   │       └── userValidator
│   │           └── user.go
│   ├── job
│   │   ├── job_payload			# 异步任务数据定义
│   │   │   └── example.go
│   │   └── job_process			# 异步任务处理器
│   │       └── example.go
│   └── task
│       └── example.go
├── main.go
├── hook.go						# 服务钩子定义
├──.env							# 本地环境配置
├──.env.development				# 开发环境配置
├──.env.production				# 生产环境配置
```

### 3. 代码生成

#### 3.1 创建项目

```shell
$ ego new projectName
```

#### 3.2 模板生成

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

### 4. 插件

```shell
# 往项目注入git commit注释规范
$ ego plugin git -t=./.git -t=commitLint
```


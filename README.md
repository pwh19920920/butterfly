# butterfly

一套简单易用的 Go 语言业务框架，基于 [gin](https://github.com/gin-gonic/gin) 封装 HTTP 服务，开箱即用、核心足够轻量。

## Features

- **HTTP 服务**：基于 gin 模块化封装，内置路由注册、中间件链与优雅关闭。
- **Config**：yaml 配置，支持命令行覆盖、按运行模式合并环境配置。
- **Logger**：基于 logrus 封装，按级别分文件、按天切割、带 traceId 透传。
- **Request / Response**：统一的请求分页与响应体结构。
- **Middleware**：内置 trace、请求日志、recover 中间件，可追加自定义中间件。

## 项目结构

```
butterfly/
├── cmd/
│   └── example/                 # 示例应用（package main）
│       └── main.go
├── configs/
│   └── config.yml               # 默认配置
├── pkg/                         # 对外公共代码
│   ├── server/                  # HTTP 服务、路由、中间件
│   │   └── middleware/          # trace / request / recover
│   ├── logger/                  # 日志封装
│   ├── response/                # 请求分页与响应体
│   ├── config/                  # 配置加载与选项
│   ├── common/                  # 公共常量
│   └── helper/                  # 工具函数
├── go.mod
└── go.sum
```

所有包位于 `github.com/pwh19920920/butterfly/pkg` 下。服务入口为 `server.StartHttpServer()`。

## Quick start

### 1. 启动服务

```go
package main

import "github.com/pwh19920920/butterfly/pkg/server"

func main() {
    server.StartHttpServer()
}
```

默认从 `configs/config.yml` 加载配置（可用 `--configFilePath` 指定其他路径）。

### 2. 编写 handler

```go
package test

import (
    "github.com/gin-gonic/gin"
    "github.com/pwh19920920/butterfly/pkg/response"
)

func test(context *gin.Context) {
    response.BuildResponseSuccess(context, "ok")
}
```

### 3. 注册路由

```go
package test

import "github.com/pwh19920920/butterfly/pkg/server"

func init() {
    route := []server.RouteInfo{
        {HttpMethod: server.HttpGet, Path: "/test", HandlerFunc: test},
    }
    server.RegisterRoute("/api", route)
}
```

在 main 包匿名引入即可生效：

```go
package main

import (
    _ "your-app/src/app/test"

    "github.com/pwh19920920/butterfly/pkg/server"
)

func main() {
    server.StartHttpServer()
}
```

可用的 HTTP 方法常量：`HttpGet`、`HttpPost`、`HttpPut`、`HttpDelete`。

完整 demo 参考：[butterfly-demo](https://github.com/pwh19920920/butterfly-demo)

## Advance

### 注册错误处理路由

```go
func init() {
    server.Register404Route(func(c *gin.Context) {
        response.Response(c, 404, "404", nil)
    })
    server.Register500Route(func(c *gin.Context) {
        response.Response(c, 500, "500", nil)
    })
}
```

另有 `Register403Route`、`Register401Route` 注册 403、401 处理器。

### 注册中间件

内置的 trace、请求日志、recover 中间件已默认启用，可用 `RegisterMiddleware` 追加自定义中间件：

```go
func init() {
    server.RegisterMiddleware(func(c *gin.Context) {
        // 自定义逻辑
        c.Next()
    })
}
```

## Config

### 命令参数

| 参数 | 说明 |
| --- | --- |
| `--configFilePath` | 配置文件地址，默认 `configs/config.yml` |
| `--server.engineMode` | 引擎模式：debug、release、test |
| `--server.serverAddr` | 启动地址，格式 `ip:端口`，仅端口则为 `:端口` |

### 配置项

```yaml
logger:
  level: 日志等级，info / warn / debug / trace / error
  logPath: 日志目录，默认 logs
  fileName: 日志文件名，默认 server
  dateFormat: 默认 2006-01-02 15:04:05.999999999
  closeConsoleOut: 关闭控制台输出，默认 false

server:
  engineMode: 引擎模式，默认 debug，可选 debug / release / test
  serverAddr: 启动地址
  serverName: 服务名称
  htmlGlobs: html 页面地址列表
  statics: 静态目录映射
  methodOverride: 是否允许路由方法覆盖，默认 false
```

加载时除主配置外，还会尝试合并 `config-<engineMode>.yml`（如 `config-release.yml`），用于按运行模式覆盖配置；该文件缺失不影响启动。

### 配置优先级

> 命令参数 > 配置参数 > 默认值

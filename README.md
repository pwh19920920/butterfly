# butterfly 快速开发框架

butterfly是一套简单易用的Go语言业务框架，整体逻辑设计简洁，支持HTTP服务、任务调度等常用业务场景模式。

# Features
1. HTTP服务：基于gin进行模块化设计，简单易用、核心足够轻量；
2. Config: 采用yaml语义化的配置文件格式，简单易用；
3. Logger: 基于logrus进行封装；
4. Request and Response：定义输入和输出数据实体格式；
5. Command Config: 支持命令启动配置、命令覆盖；

# Quick start

#### 1. 基本框架

```
import "github.com/pwh19920920/butterfly"

package main
func main() {
    butterfly.Run()
}
```

#### 2. 编写控制层
```
package test

import "github.com/gin-gonic/gin"
import "github.com/pwh19920920/butterfly/response"
import "github.com/pwh19920920/butterfly/server"

func test(context *gin.Context) {
    response.BuildResponseBadRequest(context, "test")
}
```

#### 3. 注册进butterfly

```
方式一、匿名引用导入
// 在test的init中注册路由
func init() {
    var route []server.RouteInfo
    route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "/", HandlerFunc: test})
    server.RegisterRoute("/test", route)
}

// main包匿名引入生效
import _ "butterfly-web/src/app/test" // 建议采用


方式二、main包初始方法导入
func init() {
    var route []server.RouteInfo
    route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "/", HandlerFunc: test})
    server.RegisterRoute("/test", route)
}
```

# Advance
#### 1. 注册404，500事件路由
```
// mian包初始化执行
func init() {
    server.Register404Route(func(context *gin.Context) {
        // 此处可以执行gin的正常输出
        response.Response(context, 404, "404", nil)
    })

    server.Register500Route(func(context *gin.Context) {
        // 此处可以执行gin的正常输出
        response.Response(context, 500, "500", nil)
    })
}
```

#### 2. 注册middleware
```
// mian包初始化执行
func init() {
    server.RegisterMiddleware(func(context *gin.Context) {
        // TODO 中间价逻辑
        // ...
        
        context.Next()
    })
}
```

# Config
#### 1. 命令参数
1. --configFilePath: 外部配置文件地址
2. --engineMode: 引擎模式：debug，release，test
3. --serverAddr: 启动地址：格式为ip地址:端口, 地址无限制则为:端口

#### 2. 配置参数
```yaml

logger:
  level: 日志等级，info, warn, debug, trace, error
  logPath: 日志目录，默认为logs目录
  fileName: 日志名称，默认为server
  dateFormat: 默认为2006-01-02 15:04:05.999999999
  closeConsoleOut: 关闭控制台输出，默认false
  
server:
  engineMode: 引擎模式，默认为debug，可选项debug，release，test
  serverAddr: 启动地址：格式为ip地址:端口, 地址无限制则为:端口
  serviceName: 服务名称
  htmlGlobs: html页面地址列表
  statics: 静态目录映射
```
#### 3. 配置优先级
> 命令参数 > 配置参数 > 默认值
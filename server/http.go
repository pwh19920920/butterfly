package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartHttpServer() {
	// 初始化引擎
	engine := initEngine()

	// 初始化srv
	srv := &http.Server{
		Addr:    getConf().ServerAddr,
		Handler: engine,
	}

	// log info
	logrus.Infof("server start for address '%s', running in engineMode '%s'", getConf().ServerAddr, getConf().EngineMode)

	// 服务启动
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithFields(logrus.Fields{"error": err}).Fatal("服务启动监听失败")
		}
	}()

	// 优雅关机
	gracefulShutdown(srv)
}

// 初始化引擎
func initEngine() *gin.Engine {
	// 设置启动模式
	gin.SetMode(getConf().EngineMode)

	// 启动
	engine := gin.New()

	// 注册中间价
	if middlewareList != nil {
		engine.Use(middlewareList...)
	}

	// 加载路由信息
	initRoute(engine)

	// 初始化静态资源，页面
	initLoadHTMLGlob(engine)
	initStatic(engine)
	return engine
}

// 初始化路由
func initRoute(engine *gin.Engine) {
	for _, routeGroup := range routeGroups {
		for _, routeInfo := range routeGroup.RouteInfos {
			logrus.Infof("http mvc register, method:%s, uri:%s", routeInfo.HttpMethod.String(), routeGroup.BasePath+routeInfo.Path)
			engine.Handle(routeInfo.HttpMethod.String(), routeGroup.BasePath+routeInfo.Path, routeInfo.HandlerFunc)
		}
	}

	// 404 页面处理器
	if routeFor404 != nil {
		engine.NoMethod(routeFor404)
		engine.NoRoute(routeFor404)
	}
}

// 加载页面
func initLoadHTMLGlob(engine *gin.Engine) {
	loadHTMLGlobs := conf.Server.HtmlGlobs
	if loadHTMLGlobs == nil {
		return
	}
	for _, item := range loadHTMLGlobs {
		engine.LoadHTMLGlob(item)
	}
}

// 静态配置
func initStatic(engine *gin.Engine) {
	statics := conf.Server.Statics
	if statics == nil {
		return
	}

	for key, val := range statics {
		logrus.Infof("http static register, relativePath:%s", key)
		engine.Static(key, val)
	}
}

// 优雅关机, 让服务器停个5s
func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("服务开始优雅关闭 -- 开始")

	// 服务暂停5s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Fatal("服务优雅关闭失败")
		return
	}

	logrus.Info("服务开始优雅关闭 -- 结束")
}

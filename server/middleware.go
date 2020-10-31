package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/server/middleware"
)

var middlewareList = []gin.HandlerFunc{
	middleware.Recover(routeFor500),
}

func RegisterMiddleware(handlers ...gin.HandlerFunc) {
	for _, handler := range handlers {
		middlewareList = append(middlewareList, handler)
	}
}

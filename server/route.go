package server

import "github.com/gin-gonic/gin"

type HttpMethod struct {
	Method string
}

var HttpGet = HttpMethod{"GET"}
var HttpPost = HttpMethod{"POST"}
var HttpPut = HttpMethod{"PUT"}
var HttpDelete = HttpMethod{"DELETE"}

var routeGroups = make([]RouteGroup, 1000)
var routeFor404 gin.HandlerFunc

type RouteGroup struct {
	BasePath   string
	RouteInfos []RouteInfo
}

type RouteInfo struct {
	HttpMethod  HttpMethod
	Path        string
	HandlerFunc gin.HandlerFunc
}

func (method *HttpMethod) String() string {
	return method.Method
}

func RegisterRoute(basePath string, routeInfos []RouteInfo) {
	routeGroups = append(routeGroups, RouteGroup{basePath, routeInfos})
}

func Register404Handler(handlerFunc gin.HandlerFunc) {
	routeFor404 = handlerFunc
}

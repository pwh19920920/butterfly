package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
)

type HttpMethod struct {
	Method string
}

var HttpGet = HttpMethod{"GET"}
var HttpPost = HttpMethod{"POST"}
var HttpPut = HttpMethod{"PUT"}
var HttpDelete = HttpMethod{"DELETE"}

var routeGroups = make([]RouteGroup, 0)
var routeFor404 gin.HandlerFunc = func(context *gin.Context) {
	response.Response(context, 404, "page or method not found", nil)
}

var routeFor500 gin.HandlerFunc = func(context *gin.Context) {
	response.Response(context, 500, "occurrence of system anomaly", nil)
}

var routeFor403 gin.HandlerFunc = func(context *gin.Context) {
	response.Response(context, 403, "sorry, you don't have permission for this path", nil)
}

var routeFor401 gin.HandlerFunc = func(context *gin.Context) {
	response.Response(context, 401, "sorry, you need login for the option", nil)
}

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

func RegisterRouteGroup(routeGroup RouteGroup) {
	routeGroups = append(routeGroups, routeGroup)
}

func Register404Route(handlerFunc gin.HandlerFunc) {
	routeFor404 = handlerFunc
}

func Register500Route(handlerFunc gin.HandlerFunc) {
	routeFor500 = handlerFunc
}

func Register403Route(handlerFunc gin.HandlerFunc) {
	routeFor403 = handlerFunc
}

func Register401Route(handlerFunc gin.HandlerFunc) {
	routeFor401 = handlerFunc
}

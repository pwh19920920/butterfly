package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"reflect"
	"runtime"
)

type HttpMethod struct {
	Method string
}

var HttpGet = HttpMethod{"GET"}
var HttpPost = HttpMethod{"POST"}
var HttpPut = HttpMethod{"PUT"}
var HttpDelete = HttpMethod{"DELETE"}

var routeGroupMaps = make(map[string]map[string]RouteInfo)
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
	// 如果存在则不用创建map, 如果不存在则每次都要创建map
	if _, ok := routeGroupMaps[basePath]; !ok {
		routeGroupMaps[basePath] = make(map[string]RouteInfo)
	}

	routeInfoMap := routeGroupMaps[basePath]
	for _, routeInfo := range routeInfos {
		if oldRouteInfo, ok := routeInfoMap[routeInfo.HttpMethod.String()+routeInfo.Path]; ok {
			if !GetConf().MethodOverride {
				// 不允许覆盖则需要报错
				consoleLogger.Panicf("不允许方法覆盖, tips： method:%s, uri:%s, code path: [%v] to [%v]", routeInfo.HttpMethod.String(), basePath+routeInfo.Path, methodCodePath(oldRouteInfo.HandlerFunc), methodCodePath(routeInfo.HandlerFunc))
			}
			consoleLogger.Infof("route override, method:%s, uri:%s, code path: [%v] to [%v]", routeInfo.HttpMethod.String(), basePath+routeInfo.Path, methodCodePath(oldRouteInfo.HandlerFunc), methodCodePath(routeInfo.HandlerFunc))
		}
		routeInfoMap[routeInfo.HttpMethod.String()+routeInfo.Path] = routeInfo
	}
	routeGroupMaps[basePath] = routeInfoMap
}

func methodCodePath(HandlerFunc gin.HandlerFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(HandlerFunc).Pointer()).Name()
}

func RegisterRouteGroup(routeGroup RouteGroup) {
	RegisterRoute(routeGroup.BasePath, routeGroup.RouteInfos)
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

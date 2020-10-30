package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"net/http"
)

var ResponseHelper = new(ResponseUtils)

type ResponseUtils struct {
}

func (t *ResponseUtils) buildResponse(code int, msg string, data interface{}) response.RespBody {
	return response.RespBody{
		Status:  code,
		Message: msg,
		Data:    data,
	}
}

func (t *ResponseUtils) Response(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, response.RespBody{
		Status:  code,
		Message: msg,
		Data:    data,
	})
}

func (t *ResponseUtils) ResponseSuccess(ctx *gin.Context, data interface{}) {
	t.Response(ctx, http.StatusOK, "OK", data)
}

func (t *ResponseUtils) ResponseBadRequest(ctx *gin.Context, message string) {
	t.Response(ctx, http.StatusBadRequest, message, nil)
}

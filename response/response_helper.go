package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenericResponse(code int, msg string, data interface{}) RespBody {
	return RespBody{
		Status:  code,
		Message: msg,
		Data:    data,
	}
}

func BuildResponseSuccess(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, "OK", data)
}

func BuildResponseBadRequest(ctx *gin.Context, message string) {
	Response(ctx, http.StatusBadRequest, message, nil)
}

func Response(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, GenericResponse(code, msg, data))
}

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

func BuildPageResponseSuccess(ctx *gin.Context, paging RequestPaging, total int64, data interface{}) {
	Response(ctx, http.StatusOK, "OK", RespPaging{
		RespBody: GenericResponse(http.StatusOK, "OK", data),
		PageSize: paging.GetPageSize(),
		Total:    total,
		Current:  paging.GetCurrent(),
	})
}

func BuildResponseSuccess(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, "OK", data)
}

func BuildResponseBadRequest(ctx *gin.Context, message string) {
	Response(ctx, http.StatusBadRequest, message, nil)
}

func Response(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, RespPaging{RespBody: GenericResponse(code, msg, data)})
}

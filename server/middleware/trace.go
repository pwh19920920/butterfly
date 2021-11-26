package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/pwh19920920/butterfly/common"
)

func Trace() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(common.TraceSpanHeaderKey, uuid.New())
		traceId := context.GetHeader(common.TraceIdHeaderKey)
		if traceId != "" {
			context.Set(common.TraceIdHeaderKey, traceId)
			context.Next()
			return
		}

		context.Set(common.TraceIdHeaderKey, uuid.New())
		context.Next()
	}
}

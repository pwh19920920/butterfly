package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"strings"
)

func Recover(recoverHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 异常处理
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				logrus.Error(DebugStack)
				recoverHandler(context)
				return
			}
		}()

		// 继续往下直走
		context.Next()
	}
}

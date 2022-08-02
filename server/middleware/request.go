package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/logger"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Request() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 数据装包
		writer := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: context.Writer}
		context.Writer = writer
		data, _ := context.GetRawData()
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		// 请求方式
		reqMethod := context.Request.Method

		// 请求路由
		requestURI := context.Request.RequestURI

		// 请求地址
		requestURL := context.FullPath()

		// 请求IP
		requestHost := context.Request.Host

		// 格式化
		requestBody := make(map[string]interface{}, 0)
		err := context.ShouldBindJSON(&requestBody)

		logger.InfoEntry(context, logrus.WithFields(logrus.Fields{
			"requestHost":   requestHost,
			"requestMethod": reqMethod,
			"requestURI":    requestURI,
			"requestURL":    requestURL,
			"requestBody":   requestBody,
			"requestQuery":  context.Request.URL.Query(),
		}))

		// 打印解析出错
		if err != nil {
			logger.Warn(context, "ShouldBindJSON method fail", err.Error())
		}

		// 执行下一步
		context.Next()

		// 执行时间
		latencyTime := time.Now().Sub(startTime)

		// 状态码
		statusCode := context.Writer.Status()

		// 响应部分
		logger.InfoEntry(context, logrus.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"latencyTime": fmt.Sprintf("%0.fms", latencyTime.Seconds()*1000),
			"response":    writer.body.String(),
		}))
	}
}

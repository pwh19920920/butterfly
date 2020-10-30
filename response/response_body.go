package response

type RespBody struct {
	Status  int         `json:"status"`  // 请求状态
	Data    interface{} `json:"data"`    // 数据体
	Message string      `json:"message"` // 消息
}

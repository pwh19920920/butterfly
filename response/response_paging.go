package response

type RespPaging struct {
	RespBody
	Total    int64 `json:"total"`
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
}

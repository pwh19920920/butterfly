package response

const maxPageSize = 100
const minPageSize = 20

type RequestPaging struct {
	PageSize int `form:"pageSize"` // 大小
	Current  int `form:"current"`  // 页码
}

// 偏移
func (req *RequestPaging) Offset() int {
	return (req.GetCurrent() - 1) * req.GetPageSize()
}

// 分页大小
func (req *RequestPaging) GetPageSize() int {
	if req.PageSize > maxPageSize {
		return maxPageSize
	}

	if req.PageSize <= minPageSize {
		return minPageSize
	}
	return req.PageSize
}

// 当前页码
func (req *RequestPaging) GetCurrent() int {
	if req.Current == 0 {
		return 1
	}
	return req.Current
}

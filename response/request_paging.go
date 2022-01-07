package response

const maxPageSize = 100
const minPageSize = 10

type RequestPaging struct {
	PageSize int `form:"pageSize"` // 大小
	Current  int `form:"current"`  // 页码
}

// Offset 偏移
func (req *RequestPaging) Offset() int {
	return (req.GetCurrent() - 1) * req.GetPageSize()
}

// GetPageSize 分页大小
func (req *RequestPaging) GetPageSize() int {
	if req.PageSize > maxPageSize {
		return maxPageSize
	}

	if req.PageSize <= minPageSize {
		return minPageSize
	}
	return req.PageSize
}

// GetCurrent 当前页码
func (req *RequestPaging) GetCurrent() int {
	if req.Current == 0 {
		return 1
	}
	return req.Current
}

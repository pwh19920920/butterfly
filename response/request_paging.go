package response

const maxPageSize = 100
const minPageSize = 20

type RequestPaging struct {
	PageSize int64 `form:"pageSize"` // 大小
	Current  int64 `form:"current"`  // 页码
}

// 偏移
func (req *RequestPaging) Offset() int64 {
	return (req.GetCurrent() - 1) * req.GetPageSize()
}

// 分页大小
func (req *RequestPaging) GetPageSize() int64 {
	if req.PageSize > maxPageSize {
		return maxPageSize
	}

	if req.PageSize <= minPageSize {
		return minPageSize
	}
	return req.PageSize
}

// 当前页码
func (req *RequestPaging) GetCurrent() int64 {
	if req.Current == 0 {
		return 1
	}
	return req.Current
}

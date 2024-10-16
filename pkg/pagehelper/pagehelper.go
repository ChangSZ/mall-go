package pagehelper

type ListData[T any] struct {
	PageNum   int   `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	TotalPage int64 `json:"totalPage"`
	Total     int64 `json:"total"`
	List      []T   `json:"list"`
}

func New[T any]() *ListData[T] {
	return &ListData[T]{}
}

func (s *ListData[T]) Set(pageNum, pageSize int, total int64, list []T) *ListData[T] {
	s.PageNum = pageNum
	s.PageSize = pageSize
	s.Total = total
	totalPage := total / int64(pageSize)
	if total%int64(pageSize) > 0 {
		totalPage += 1
	}
	s.TotalPage = totalPage
	s.List = list
	return s
}

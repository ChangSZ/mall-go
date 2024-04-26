package dto

type UriID struct {
	Id int64 `uri:"id" binding:"required"` // 用户ID
}

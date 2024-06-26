package member_read_history

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/member_read_history"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建浏览记录
	// @Tags MemberCollectionController
	// @Router /member/readHistory/create [post]
	Create(*gin.Context)

	// Delete 删除浏览记录
	// @Tags MemberCollectionController
	// @Router /member/readHistory/delete [post]
	Delete(*gin.Context)

	// List 分页获取浏览记录
	// @Tags MemberCollectionController
	// @Router /member/readHistory/list [get]
	List(*gin.Context)

	// Clear 清空浏览记录
	// @Tags MemberCollectionController
	// @Router /member/readHistory/clear [post]
	Clear(*gin.Context)
}

type handler struct {
	service member_read_history.Service
}

func New() Handler {
	return &handler{
		service: member_read_history.New(),
	}
}

func (h *handler) i() {}

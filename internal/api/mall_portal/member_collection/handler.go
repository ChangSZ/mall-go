package member_collection

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/member_collection"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Add 添加商品收藏
	// @Tags MemberCollectionController
	// @Router /member/productCollection/add [post]
	Add(*gin.Context)

	// Delete 删除商品收藏
	// @Tags MemberCollectionController
	// @Router /member/productCollection/delete [post]
	Delete(*gin.Context)

	// List 显示当前用户商品收藏列表
	// @Tags MemberCollectionController
	// @Router /member/productCollection/list [get]
	List(*gin.Context)

	// Detail 显示商品收藏详情
	// @Tags MemberCollectionController
	// @Router /member/productCollection/detail [get]
	Detail(*gin.Context)

	// Clear 清空当前用户商品收藏列表
	// @Tags MemberCollectionController
	// @Router /member/productCollection/clear [post]
	Clear(*gin.Context)
}

type handler struct {
	service member_collection.Service
}

func New() Handler {
	return &handler{
		service: member_collection.New(),
	}
}

func (h *handler) i() {}

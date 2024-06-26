package member_attention

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/member_attention"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Add 添加品牌关注
	// @Tags MemberAttentionController
	// @Router /member/attention/add [post]
	Add(*gin.Context)

	// Delete 取消品牌关注
	// @Tags MemberAttentionController
	// @Router /member/attention/delete [post]
	Delete(*gin.Context)

	// List 分页查询当前用户品牌关注列表
	// @Tags MemberAttentionController
	// @Router /member/attention/list [get]
	List(*gin.Context)

	// Detail 根据品牌ID获取品牌关注详情
	// @Tags MemberAttentionController
	// @Router /member/attention/detail [get]
	Detail(*gin.Context)

	// Clear 清空当前用户品牌关注列表
	// @Tags MemberAttentionController
	// @Router /member/attention/clear [post]
	Clear(*gin.Context)
}

type handler struct {
	service member_attention.Service
}

func New() Handler {
	return &handler{
		service: member_attention.New(),
	}
}

func (h *handler) i() {}

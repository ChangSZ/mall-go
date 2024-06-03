package ums_member_receive_address

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member_receive_address"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Add 添加收货地址
	// @Tags UmsMemberReceiveAddressController
	// @Router /member/address/add [post]
	Add(*gin.Context)

	// Delete 删除收货地址
	// @Tags UmsMemberReceiveAddressController
	// @Router /member/address/delete/{id} [post]
	Delete(*gin.Context)

	// Update 修改收货地址
	// @Tags UmsMemberReceiveAddressController
	// @Router /member/address/update/{id} [post]
	Update(*gin.Context)

	// List 获取所有收货地址
	// @Tags UmsMemberReceiveAddressController
	// @Router /member/address/list [get]
	List(*gin.Context)

	// GetItem 获取收货地址详情
	// @Tags UmsMemberReceiveAddressController
	// @Router /member/address/{id} [get]
	GetItem(*gin.Context)
}

type handler struct {
	service ums_member_receive_address.Service
}

func New() Handler {
	return &handler{
		service: ums_member_receive_address.New(),
	}
}

func (h *handler) i() {}

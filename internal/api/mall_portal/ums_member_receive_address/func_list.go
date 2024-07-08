package ums_member_receive_address

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct {
	List []dto.UmsMemberReceiveAddress `json:",inline"`
}

// List 获取所有收货地址
// @Summary 获取所有收货地址
// @Description 获取所有收货地址
// @Tags UmsMemberReceiveAddressController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /member/address/list [get]
func (h *handler) List(ctx *gin.Context) {
	_ = new(listRequest)
	res := new(listResponse)
	list, err := h.service.List(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

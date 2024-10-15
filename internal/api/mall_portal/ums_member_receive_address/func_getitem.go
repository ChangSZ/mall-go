package ums_member_receive_address

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getItemRequest struct{}

type getItemResponse struct {
	dto.UmsMemberReceiveAddress `json:",inline"`
}

// GetItem 获取收货地址详情
// @Summary 获取收货地址详情
// @Description 获取收货地址详情
// @Tags UmsMemberReceiveAddressController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /member/address/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	_ = new(getItemRequest)
	res := new(getItemResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	item, err := h.service.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.UmsMemberReceiveAddress = *item
	api.Success(ctx, res)
}

package ums_member_receive_address

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	dto.UmsMemberReceiveAddress `json:",inline"`
}

type addResponse struct {
	Count int64 `json:",inline"`
}

// Add 添加收货地址
// @Summary 添加收货地址
// @Description 添加收货地址
// @Tags UmsMemberReceiveAddressController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body addRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /member/address/add [post]
func (h *handler) Add(ctx *gin.Context) {
	req := new(addRequest)
	res := new(addResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.Add(ctx, req.UmsMemberReceiveAddress)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

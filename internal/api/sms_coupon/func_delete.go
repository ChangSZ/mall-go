package sms_coupon

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct {
	Count int64 `json:",inline"`
}

// Delete 删除优惠券
// @Summary 删除优惠券
// @Description 删除优惠券
// @Tags SmsCouponController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /coupon/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {
	_ = new(deleteRequest)
	res := new(deleteResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.smsCouponService.Delete(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "删除个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
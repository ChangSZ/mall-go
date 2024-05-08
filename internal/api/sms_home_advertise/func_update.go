package sms_home_advertise

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	dto.SmsHomeAdvertise `json:",inline"`
}

type updateResponse struct {
	Count int64 `json:",inline"`
}

// Update 修改广告
// @Summary 修改广告
// @Description 修改广告
// @Tags SmsHomeAdvertiseController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /home/advertise/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	req := new(updateRequest)
	res := new(updateResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.smsHomeAdvertiseService.Update(ctx, uri.Id, req.SmsHomeAdvertise)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

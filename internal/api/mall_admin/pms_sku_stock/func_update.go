package pms_sku_stock

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateRequest struct{}

type updateResponse struct {
	Count int64 `json:",inline"`
}

// Update 批量更新sku库存信息
// @Summary 批量更新sku库存信息
// @Description 批量更新sku库存信息
// @Tags PmsSkuStockController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body []dto.PmsSkuStock true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /sku/update/{pid} [post]
func (h *handler) Update(ctx *gin.Context) {
	_ = new(updateRequest)
	req := make([]dto.PmsSkuStock, 0)
	res := new(updateResponse)
	uri := new(dto.PmsPidUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.Update(ctx, uri.Pid, req)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

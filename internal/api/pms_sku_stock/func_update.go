package pms_sku_stock

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
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
// @Param Request body []dto.PmsSkuStockUpdateParam true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /sku/update/{pid} [post]
func (h *handler) Update(ctx *gin.Context) {
	req := make([]dto.PmsSkuStockUpdateParam, 0)
	res := new(updateResponse)
	uri := new(dto.PmsPidUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.pmsSkuStockService.Update(ctx, uri.Pid, req)
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

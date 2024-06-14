package pms_portal_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type detailRequest struct {
	BrandId int64 `uri:"brandId" binding:"required"`
}

type detailResponse struct {
	dto.PmsBrand `json:",inline"`
}

// Detail 获取品牌详情
// @Summary 获取品牌详情
// @Description 获取品牌详情
// @Tags PmsPortalBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} code.Success{data=detailResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/detail/{brandId} [get]
func (h *handler) Detail(ctx *gin.Context) {
	req := new(detailRequest)
	res := new(detailResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.Detail(ctx, req.BrandId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PmsBrand = *data
	api.Success(ctx, res)
}

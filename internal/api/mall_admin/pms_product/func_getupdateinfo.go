package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getUpdateInfoRequest struct{}

type getUpdateInfoResponse struct {
	dto.PmsProductResult `json:",inline"`
}

// GetUpdateInfo 根据商品id获取商品编辑信息
// @Summary 根据商品id获取商品编辑信息
// @Description 根据商品id获取商品编辑信息
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getUpdateInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=getUpdateInfoResponse}
// @Failure 400 {object} code.Failure
// @Router /product/updateInfo/{id} [post]
func (h *handler) GetUpdateInfo(ctx *gin.Context) {
	_ = new(getUpdateInfoRequest)
	res := new(getUpdateInfoResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	info, err := h.service.GetUpdateInfo(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PmsProductResult = *info
	api.Success(ctx, res)
}

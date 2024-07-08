package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct {
	dto.PmsBrand `json:",inline"`
}

// GetItem 根据编号查询品牌信息
// @Summary 根据编号查询品牌信息
// @Description 根据编号查询品牌信息
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	_ = new(getItemRequest)
	res := new(getItemResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	item, err := h.service.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PmsBrand = *item
	api.Success(ctx, res)
}

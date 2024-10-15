package pms_sku_stock

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getListRequest struct {
	Keyword string `form:"keyword"`
}

type getListResponse struct {
	List []dto.PmsSkuStock `json:",inline"`
}

// GetList 根据商品ID及sku编码模糊搜索sku库存
// @Summary 根据商品ID及sku编码模糊搜索sku库存
// @Description 根据商品ID及sku编码模糊搜索sku库存
// @Tags PmsSkuStockController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsSkuStock}
// @Failure 400 {object} code.Failure
// @Router /sku/{pid} [get]
func (h *handler) GetList(ctx *gin.Context) {
	req := new(getListRequest)
	res := new(getListResponse)
	uri := new(dto.PmsPidUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.ListAll(ctx, uri.Pid, req.Keyword)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

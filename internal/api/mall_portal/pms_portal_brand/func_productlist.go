package pms_portal_brand

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type productListRequest struct {
	BrandId  int64 `form:"brandId" binding:"required"`
	PageNum  int   `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int   `form:"pageSize,default=6" binding:"omitempty"`
}

type productListResponse struct {
	*pagehelper.ListData[dto.PmsProduct] `json:",inline"`
}

// ProductList 分页获取品牌相关商品
// @Summary 分页获取品牌相关商品
// @Description 分页获取品牌相关商品
// @Tags PmsPortalBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body productListRequest true "请求信息"
// @Success 200 {object} code.Success{data=productListResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/productList [get]
func (h *handler) ProductList(ctx *gin.Context) {
	req := new(productListRequest)
	res := new(productListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.ProductList(ctx, req.BrandId, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}

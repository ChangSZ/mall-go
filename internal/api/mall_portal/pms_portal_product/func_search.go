package pms_portal_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type searchRequest struct {
	Keyword           string `form:"keyword" binding:"omitempty"`
	BrandId           int64  `form:"brandId" binding:"omitempty"`
	ProductCategoryId int64  `form:"productCategoryId" binding:"omitempty"`
	PageNum           int    `form:"pageNum,default=1" binding:"omitempty"`
	PageSize          int    `form:"pageSize,default=5" binding:"omitempty"`
	Sort              int    `form:"sort,default=0" binding:"omitempty"`
}

type searchResponse struct {
	*pagehelper.ListData[dto.PmsProduct] `json:",inline"`
}

// Search 综合搜索、筛选、排序
// @Summary 综合搜索、筛选、排序
// @Description 综合搜索、筛选、排序
// @Tags PmsPortalProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body searchRequest true "请求信息"
// @Success 200 {object} code.Success{data=searchResponse}
// @Failure 400 {object} code.Failure
// @Router /product/search [get]
func (h *handler) Search(ctx *gin.Context) {
	req := new(searchRequest)
	res := new(searchResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.Search(ctx,
		req.Keyword, req.BrandId, req.ProductCategoryId, req.PageNum, req.PageSize, req.Sort)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}

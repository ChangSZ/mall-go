package home

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getProductCateListRequest struct {
	ParentId int64 `uri:"parentId"`
}

type getProductCateListResponse struct {
	List []dto.PmsProductCategory `json:",inline"`
}

// GetProductCateList 获取首页商品分类
// @Summary 获取首页商品分类
// @Description 获取首页商品分类
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getProductCateListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProductCategory}
// @Failure 400 {object} code.Failure
// @Router /home/productCateList/{parentId} [get]
func (h *handler) GetProductCateList(ctx *gin.Context) {
	req := new(getProductCateListRequest)
	res := new(getProductCateListResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.GetProductCateList(ctx, req.ParentId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = data
	api.Success(ctx, res.List)
}

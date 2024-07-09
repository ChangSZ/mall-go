package pms_product_attr

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getAttrInfoRequest struct{}

type getAttrInfoResponse struct {
	List []dto.PmsProductAttrInfo `json:",inline"`
}

// GetAttrInfo 根据商品分类的id获取商品属性及属性分类
// @Summary 根据商品分类的id获取商品属性及属性分类
// @Description 根据商品分类的id获取商品属性及属性分类
// @Tags PmsProductAttributeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getAttrInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProductAttrInfo}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/attrInfo/{productCategoryId} [get]
func (h *handler) GetAttrInfo(ctx *gin.Context) {
	_ = new(getAttrInfoRequest)
	res := new(getAttrInfoResponse)
	uri := new(dto.PmsProductCateIdUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	info, err := h.service.GetProductAttrInfo(ctx, uri.ProductCategoryId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = info
	api.Success(ctx, res.List)
}

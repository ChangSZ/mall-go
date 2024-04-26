package pms_product_attr

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getAttrInfoRequest struct{}

type getAttrInfoResponse struct{}

// GetAttrInfo 根据商品分类的id获取商品属性及属性分类
// @Summary 根据商品分类的id获取商品属性及属性分类
// @Description 根据商品分类的id获取商品属性及属性分类
// @Tags PmsProductAttributeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getAttrInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=getAttrInfoResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/attrInfo/{productCategoryId} [get]
func (h *handler) GetAttrInfo(ctx *gin.Context) {
	api.Success(ctx, nil)
}

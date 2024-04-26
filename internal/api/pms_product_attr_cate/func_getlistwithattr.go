package pms_product_attr_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getListWithAttrRequest struct{}

type getListWithAttrResponse struct{}

// GetListWithAttr 获取所有商品属性分类及其下属性
// @Summary 获取所有商品属性分类及其下属性
// @Description 获取所有商品属性分类及其下属性
// @Tags PmsProductAttributeCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getListWithAttrRequest true "请求信息"
// @Success 200 {object} code.Success{data=getListWithAttrResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/category/list/withAttr [get]
func (h *handler) GetListWithAttr(ctx *gin.Context) {
	api.Success(ctx, nil)
}

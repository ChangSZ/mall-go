package pms_product_attr_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type getListWithAttrRequest struct{}

type getListWithAttrResponse struct {
	List []dto.PmsProductAttrCateItem `json:",inline"`
}

// GetListWithAttr 获取所有商品属性分类及其下属性
// @Summary 获取所有商品属性分类及其下属性
// @Description 获取所有商品属性分类及其下属性
// @Tags PmsProductAttributeCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getListWithAttrRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProductAttrCateItem}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/category/list/withAttr [get]
func (h *handler) GetListWithAttr(ctx *gin.Context) {
	_ = new(getListWithAttrRequest)
	res := new(getListWithAttrResponse)
	list, err := h.pmsProductAttrCateService.ListWithAttr(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

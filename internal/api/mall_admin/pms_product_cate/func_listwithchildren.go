package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listWithChildrenRequest struct{}

type listWithChildrenResponse struct {
	List []dto.PmsProductCategoryWithChildrenItem `json:",inline"`
}

// ListWithChildren 查询所有一级分类及子分类
// @Summary 查询所有一级分类及子分类
// @Description 查询所有一级分类及子分类
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listWithChildrenRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProductCategoryWithChildrenItem}
// @Failure 400 {object} code.Failure
// @Router /productCategory/list/withChildren [get]
func (h *handler) ListWithChildren(ctx *gin.Context) {
	_ = new(listWithChildrenRequest)
	res := new(listWithChildrenResponse)
	list, err := h.service.ListWithChildren(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

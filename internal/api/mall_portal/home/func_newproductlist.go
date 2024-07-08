package home

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type newProductListRequest struct {
	PageNum  int `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int `form:"pageSize,default=6" binding:"omitempty"`
}

type newProductListResponse struct {
	List []dto.PmsProduct `json:",inline"`
}

// NewProductList 分页获取新品推荐商品
// @Summary 分页获取新品推荐商品
// @Description 分页获取新品推荐商品
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body newProductListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProduct}
// @Failure 400 {object} code.Failure
// @Router /home/newProductList [get]
func (h *handler) NewProductList(ctx *gin.Context) {
	req := new(newProductListRequest)
	res := new(newProductListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.NewProductList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

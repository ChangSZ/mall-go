package home

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type hotProductListRequest struct {
	PageNum  int `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int `form:"pageSize,default=6" binding:"omitempty"`
}

type hotProductListResponse struct {
	List []dto.PmsProduct `json:",inline"`
}

// HotProductList 分页获取人气推荐商品
// @Summary 分页获取人气推荐商品
// @Description 分页获取人气推荐商品
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body hotProductListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProduct}
// @Failure 400 {object} code.Failure
// @Router /home/hotProductList [get]
func (h *handler) HotProductList(ctx *gin.Context) {
	req := new(hotProductListRequest)
	res := new(hotProductListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.HotProductList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

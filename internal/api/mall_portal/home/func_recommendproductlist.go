package home

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type recommendProductListRequest struct {
	PageNum  int `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int `form:"pageSize,default=4" binding:"omitempty"`
}

type recommendProductListResponse struct {
	List []dto.PmsProduct `json:",inline"`
}

// RecommendProductList 分页获取推荐商品
// @Summary 分页获取推荐商品
// @Description 分页获取推荐商品
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body recommendProductListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProduct}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProductList [get]
func (h *handler) RecommendProductList(ctx *gin.Context) {
	req := new(recommendProductListRequest)
	res := new(recommendProductListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.RecommendProductList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

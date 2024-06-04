package pms_portal_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type recommendListRequest struct {
	PageNum  int `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int `form:"pageSize,default=5" binding:"omitempty"`
}

type recommendListResponse struct {
	List []dto.PmsBrand `json:",inline"`
}

// RecommendList 分页获取推荐品牌
// @Summary 分页获取推荐品牌
// @Description 分页获取推荐品牌
// @Tags PmsPortalBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body recommendListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsBrand}
// @Failure 400 {object} code.Failure
// @Router /brand/recommendList [get]
func (h *handler) RecommendList(ctx *gin.Context) {
	req := new(recommendListRequest)
	res := new(recommendListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.RecommendList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

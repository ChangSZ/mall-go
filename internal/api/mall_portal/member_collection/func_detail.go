package member_collection

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type detailRequest struct {
	ProductId int64 `form:"productId" binding:"required"`
}

type detailResponse struct {
	*dto.MemberProductCollection `json:",inline"`
}

// Detail 显示商品收藏详情
// @Summary 显示商品收藏详情
// @Description 显示商品收藏详情
// @Tags MemberCollectionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} code.Success{data=detailResponse}
// @Failure 400 {object} code.Failure
// @Router /member/productCollection/detail [get]
func (h *handler) Detail(ctx *gin.Context) {
	req := new(detailRequest)
	res := new(detailResponse)
	if err := ctx.ShouldBindQuery(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.Detail(ctx, req.ProductId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.MemberProductCollection = data
	api.Success(ctx, res)
}

package member_attention

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type detailRequest struct {
	BrandId int64 `form:"brandId" binding:"required"`
}

type detailResponse struct {
	*dto.MemberBrandAttention `json:",inline"`
}

// Detail 根据品牌ID获取品牌关注详情
// @Summary 根据品牌ID获取品牌关注详情
// @Description 根据品牌ID获取品牌关注详情
// @Tags MemberAttentionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} code.Success{data=detailResponse}
// @Failure 400 {object} code.Failure
// @Router /member/attention/detail [get]
func (h *handler) Detail(ctx *gin.Context) {
	req := new(detailRequest)
	res := new(detailResponse)
	if err := ctx.ShouldBindQuery(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.Detail(ctx, req.BrandId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.MemberBrandAttention = data
	api.Success(ctx, res)
}

package sms_home_recommend_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct {
	Count int64 `json:",inline"`
}

// Create 添加首页推荐
// @Summary 添加首页推荐
// @Description 添加首页推荐
// @Tags SmsHomeRecommendProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProduct [post]
func (h *handler) Create(ctx *gin.Context) {
	_ = new(createRequest)
	req := make([]dto.SmsHomeRecommendProduct, 0)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.smsHomeRecommendProductService.Create(ctx, req)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "创建个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

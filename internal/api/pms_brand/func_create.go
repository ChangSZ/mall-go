package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	dto.PmsBrandParam `json:",inline"`
}

type createResponse struct {
	Count int64 `json:",inline"`
}

// Create 添加品牌
// @Summary 添加品牌
// @Description 添加品牌
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /brand/create [post]
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.pmsBrandService.Create(ctx, req.PmsBrandParam)
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
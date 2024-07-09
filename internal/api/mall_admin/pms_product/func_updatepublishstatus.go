package pms_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updatePublishStatusRequest struct {
	Ids           []int64 `form:"ids"`
	PublishStatus int32   `form:"publishStatus"`
}

type updatePublishStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdatePublishStatus 批量上下架商品
// @Summary 批量上下架商品
// @Description 批量上下架商品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePublishStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /product/update/publishStatus [post]
func (h *handler) UpdatePublishStatus(ctx *gin.Context) {
	req := new(updatePublishStatusRequest)
	res := new(updatePublishStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdatePublishStatus(ctx, req.Ids, req.PublishStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

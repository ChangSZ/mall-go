package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateNewStatusRequest struct {
	Ids       []int64 `form:"ids"`
	NewStatus int32   `form:"newStatus"`
}

type updateNewStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateNewStatus 批量设为新品
// @Summary 批量设为新品
// @Description 批量设为新品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateNewStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /product/update/newStatus [post]
func (h *handler) UpdateNewStatus(ctx *gin.Context) {
	req := new(updateNewStatusRequest)
	res := new(updateNewStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.pmsProductService.UpdateNewStatus(ctx, req.Ids, req.NewStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新NewStatus个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

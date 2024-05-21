package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateVerifyStatusRequest struct {
	Ids          []int64 `form:"ids"`
	VerifyStatus int32   `form:"verifyStatus"`
	Detail       string  `form:"detail"`
}

type updateVerifyStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateVerifyStatus 批量修改审核状态
// @Summary 批量修改审核状态
// @Description 批量修改审核状态
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateVerifyStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /product/update/verifyStatus [post]
func (h *handler) UpdateVerifyStatus(ctx *gin.Context) {
	req := new(updateVerifyStatusRequest)
	res := new(updateVerifyStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateVerifyStatus(ctx, req.Ids, req.VerifyStatus, req.Detail)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

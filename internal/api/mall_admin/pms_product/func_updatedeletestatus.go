package pms_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateDeleteStatusRequest struct {
	Ids          []int64 `form:"ids"`
	DeleteStatus int32   `form:"deleteStatus"`
}

type updateDeleteStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateDeleteStatus 批量修改删除状态
// @Summary 批量修改删除状态
// @Description 批量修改删除状态
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateDeleteStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /product/update/deleteStatus [post]
func (h *handler) UpdateDeleteStatus(ctx *gin.Context) {
	req := new(updateDeleteStatusRequest)
	res := new(updateDeleteStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.UpdateDeleteStatus(ctx, req.Ids, req.DeleteStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

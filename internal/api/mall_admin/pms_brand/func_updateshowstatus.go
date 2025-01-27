package pms_brand

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateShowStatusRequest struct {
	Ids        []int64 `form:"ids"`
	ShowStatus int32   `form:"showStatus"`
}

type updateShowStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateShowStatus 批量更新显示状态
// @Summary 批量更新显示状态
// @Description 批量更新显示状态
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateShowStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /brand/update/showStatus [post]
func (h *handler) UpdateShowStatus(ctx *gin.Context) {
	req := new(updateShowStatusRequest)
	res := new(updateShowStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.UpdateShowStatus(ctx, req.Ids, req.ShowStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

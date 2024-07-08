package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type updateShowStatusRequest struct {
	Ids        []int64 `form:"ids"`
	ShowStatus int32   `form:"showStatus"`
}

type updateShowStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateShowStatus 修改显示状态
// @Summary 修改显示状态
// @Description 修改显示状态
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateShowStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /productCategory/update/showStatus [post]
func (h *handler) UpdateShowStatus(ctx *gin.Context) {
	req := new(updateShowStatusRequest)
	res := new(updateShowStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
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

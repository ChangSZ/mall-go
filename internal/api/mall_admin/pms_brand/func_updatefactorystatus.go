package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type updateFactoryStatusRequest struct {
	Ids           []int64 `form:"ids"`
	FactoryStatus int32   `form:"factoryStatus"`
}

type updateFactoryStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateFactoryStatus 批量更新厂家制造商状态
// @Summary 批量更新厂家制造商状态
// @Description 批量更新厂家制造商状态
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateFactoryStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /brand/update/factoryStatus [post]
func (h *handler) UpdateFactoryStatus(ctx *gin.Context) {
	req := new(updateFactoryStatusRequest)
	res := new(updateFactoryStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateFactoryStatus(ctx, req.Ids, req.FactoryStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

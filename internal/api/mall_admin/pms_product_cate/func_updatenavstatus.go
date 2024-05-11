package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateNavStatusRequest struct {
	Ids       []int64 `form:"ids"`
	NavStatus int32   `form:"navStatus"`
}

type updateNavStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateNavStatus 修改导航栏显示状态
// @Summary 修改导航栏显示状态
// @Description 修改导航栏显示状态
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateNavStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /productCategory/update/navStatus [post]
func (h *handler) UpdateNavStatus(ctx *gin.Context) {
	req := new(updateNavStatusRequest)
	res := new(updateNavStatusResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateNavStatus(ctx, req.Ids, req.NavStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新NavStatus个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

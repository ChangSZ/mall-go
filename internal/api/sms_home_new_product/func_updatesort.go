package sms_home_new_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateSortRequest struct {
	Sort int32 `form:"sort"`
}

type updateSortResponse struct {
	Count int64 `json:",inline"`
}

// UpdateSort 修改首页新品排序
// @Summary 修改首页新品排序
// @Description 修改首页新品排序
// @Tags SmsHomeNewProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateSortRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /home/newProduct/update/sort/{id} [post]
func (h *handler) UpdateSort(ctx *gin.Context) {
	req := new(updateSortRequest)
	res := new(updateSortResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateSort(ctx, uri.Id, req.Sort)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新排序状态个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

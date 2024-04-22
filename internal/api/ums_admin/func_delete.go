package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct {
	Count int64
}

// Delete 删除指定用户信息
// @Summary 删除指定用户信息
// @Description 删除指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse.Count}
// @Failure 400 {object} code.Failure
// @Router /admin/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {
	_ = new(deleteRequest)
	res := new(deleteResponse)
	uri := new(UmsAdminUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	cnt, err := h.umsAdminService.Delete(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "未找到用户数据, 删除失败")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

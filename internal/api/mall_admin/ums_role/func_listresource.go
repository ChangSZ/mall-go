package ums_role

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listResourceRequest struct{}

type listResourceResponse struct {
	List []dto.UmsResource
}

// ListResource 获取角色相关资源
// @Summary 获取角色相关资源
// @Description 获取角色相关资源
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listResourceRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResourceResponse}
// @Failure 400 {object} code.Failure
// @Router /role/listResource/{roleId} [get]
func (h *handler) ListResource(ctx *gin.Context) {
	_ = new(listResourceRequest)
	res := new(listResourceResponse)
	uri := new(dto.UmsRoleIdUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.ListResource(ctx, uri.RoleId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}

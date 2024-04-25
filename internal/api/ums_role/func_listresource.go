package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/api/ums_resource"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listResourceRequest struct{}

type listResourceResponse struct {
	List []ums_resource.UmsResource
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
	uri := new(UmsRoleIdUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.umsRoleService.ListResource(ctx, uri.RoleId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	listData := make([]ums_resource.UmsResource, 0, len(list))
	for _, v := range list {
		listData = append(listData, ums_resource.UmsResource{
			Id:          v.Id,
			CreateTime:  v.CreateTime,
			Name:        v.Name,
			Url:         v.Url,
			Description: v.Description,
			CategoryId:  v.CategoryId,
		})
	}
	res.List = listData
	api.Success(ctx, res.List)
}

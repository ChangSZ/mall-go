package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	UmsRoleParam `json:",inline"`
}

type createResponse struct {
	Count int64 `json:"count"`
}

// Create 添加角色
// @Summary 添加角色
// @Description 添加角色
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /role/create [post]
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	data := &ums_role.UmsRole{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	cnt, err := h.umsRoleService.Create(ctx, data)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "创建个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}

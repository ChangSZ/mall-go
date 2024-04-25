package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct {
	List []UmsRole `json:",inline"`
}

// ListAll 获取所有角色
// @Summary 获取所有角色
// @Description 获取所有角色
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]UmsRole}
// @Failure 400 {object} code.Failure
// @Router /role/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {
	_ = new(listAllRequest)
	res := new(listAllResponse)
	list, err := h.umsRoleService.ListAll(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	listData := make([]UmsRole, 0, len(list))
	for _, v := range list {
		listData = append(listData, UmsRole{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			AdminCount:  v.AdminCount,
			CreateTime:  v.CreateTime,
			Status:      v.Status,
			Sort:        v.Sort,
		})
	}
	res.List = listData
	api.Success(ctx, res.List)
}

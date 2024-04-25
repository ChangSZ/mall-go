package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Keyword  string `form:"keyword" binding:"omitempty"`
	PageSize int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum  int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int       `json:"pageNum"`
	PageSize  int       `json:"pageSize"`
	TotalPage int64     `json:"totalPage"`
	Total     int64     `json:"total"`
	List      []UmsRole `json:"list"`
}

// List 根据角色名称分页获取角色列表
// @Summary 根据角色名称分页获取角色列表
// @Description 根据角色名称分页获取角色列表
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /role/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, total, err := h.umsRoleService.List(ctx, req.Keyword, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	totalPage := total / int64(req.PageSize)
	if total%int64(req.PageSize) > 0 {
		totalPage += 1
	}
	res.TotalPage = totalPage
	res.Total = total
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
	api.Success(ctx, res)
}

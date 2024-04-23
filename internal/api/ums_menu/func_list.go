package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	PageSize int `form:"pageSize" binding:"omitempty"`
	PageNum  int `form:"pageNum" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int       `json:"pageNum"`
	PageSize  int       `json:"pageSize"`
	TotalPage int64     `json:"totalPage"`
	Total     int64     `json:"total"`
	List      []UmsMenu `json:"list"`
}

// List 分页查询后台菜单
// @Summary 分页查询后台菜单
// @Description 分页查询后台菜单
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request formData listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /menu/list/{parentId} [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	uri := new(UmsMenuListUri)
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
	if req.PageSize == 0 {
		req.PageSize = 5
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	list, total, err := h.umsMenuService.List(ctx, uri.ParentId, req.PageSize, req.PageNum)
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
	listData := make([]UmsMenu, 0, len(list))
	for _, v := range list {
		listData = append(listData, UmsMenu{
			Id:         v.Id,
			ParentId:   v.ParentId,
			CreateTime: v.CreateTime,
			Title:      v.Title,
			Level:      v.Level,
			Sort:       v.Sort,
			Name:       v.Name,
			Icon:       v.Icon,
			Hidden:     v.Hidden,
		})
	}
	res.List = listData
	api.Success(ctx, res)
}

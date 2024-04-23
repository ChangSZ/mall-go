package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	UmsMenuParam `json:",inline"`
}

type createResponse struct {
	Count int64 `json:"count"`
}

// Create 添加后台菜单
// @Summary 添加后台菜单
// @Description 添加后台菜单
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /menu/create [post]
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	data := &ums_menu.UmsMenu{
		ParentId: req.ParentId,
		Title:    req.Title,
		Level:    req.Level,
		Sort:     req.Sort,
		Name:     req.Name,
		Icon:     req.Icon,
		Hidden:   req.Hidden,
	}
	cnt, err := h.umsMenuService.Create(ctx, data)
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

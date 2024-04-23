package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getRequest struct{}

type getResponse struct {
	UmsMenu `json:",inline"`
}

// Get 根据ID获取菜单详情
// @Summary 根据ID获取菜单详情
// @Description 根据ID获取菜单详情
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} code.Success{data=getResponse}
// @Failure 400 {object} code.Failure
// @Router /menu/{id} [get]
func (h *handler) Get(ctx *gin.Context) {
	_ = new(getRequest)
	res := new(getResponse)
	uri := new(UmsMenuUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	item, err := h.umsMenuService.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.UmsMenu = UmsMenu{
		Id:         item.Id,
		ParentId:   item.ParentId,
		CreateTime: item.CreateTime,
		Title:      item.Title,
		Level:      item.Level,
		Sort:       item.Sort,
		Name:       item.Name,
		Icon:       item.Icon,
		Hidden:     item.Hidden,
	}
	api.Success(ctx, res)
}

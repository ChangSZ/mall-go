package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createActionRequest struct {
	Id     string `form:"id"`     // HashID
	Method string `form:"method"` // 请求方法
	API    string `form:"api"`    // 请求地址
}

type createActionResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// CreateAction 创建功能权限
// @Summary 创建功能权限
// @Description 创建功能权限
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "HashID"
// @Param method formData string true "请求方法"
// @Param api formData string true "请求地址"
// @Success 200 {object} createActionResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu_action [post]
// @Security LoginToken
func (h *handler) CreateAction(ctx *gin.Context) {
	req := new(createActionRequest)
	res := new(createActionResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	ids, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	id := int32(ids[0])

	searchOneData := new(menu.SearchOneData)
	searchOneData.Id = id
	menuInfo, err := h.menuService.Detail(ctx, searchOneData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuDetailError, err)
		return
	}

	createActionData := new(menu.CreateMenuActionData)
	createActionData.MenuId = menuInfo.Id
	createActionData.Method = req.Method
	createActionData.API = req.API

	createId, err := h.menuService.CreateAction(ctx, createActionData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuCreateActionError, err)
		return
	}

	res.Id = createId
	api.ResponseOK(ctx, res)
}

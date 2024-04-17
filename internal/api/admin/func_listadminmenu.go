package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listAdminMenuRequest struct {
	Id string `uri:"id"` // HashID
}

type listAdminMenuResponse struct {
	List     []admin.ListMenuData `json:"list"`
	UserName string               `json:"username"`
}

// ListAdminMenu 菜单授权列表
// @Summary 菜单授权列表
// @Description 菜单授权列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} listAdminMenuResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/menu/{id} [get]
// @Security LoginToken
func (h *handler) ListAdminMenu(ctx *gin.Context) {
	req := new(listAdminMenuRequest)
	res := new(listAdminMenuResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
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

	searchOneData := new(admin.SearchOneData)
	searchOneData.Id = int32(ids[0])
	searchOneData.IsUsed = 1

	info, err := h.adminService.Detail(ctx, searchOneData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminMenuListError, err)
		return
	}

	res.UserName = info.Username

	searchData := new(admin.SearchListMenuData)
	searchData.AdminId = int32(ids[0])

	listData, err := h.adminService.ListMenu(ctx, searchData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminMenuListError, err)
		return
	}

	res.List = listData
	api.ResponseOK(ctx, res)
}

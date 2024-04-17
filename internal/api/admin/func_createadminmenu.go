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

type createAdminMenuRequest struct {
	Id      string `form:"id"`      // HashID
	Actions string `form:"actions"` // 功能权限ID,多个用,分割
}

type createAdminMenuResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// CreateAdminMenu 提交菜单授权
// @Summary 提交菜单授权
// @Description 提交菜单授权
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "Hashid"
// @Param actions formData string true "功能权限ID,多个用,分割"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/menu [post]
// @Security LoginToken
func (h *handler) CreateAdminMenu(ctx *gin.Context) {
	req := new(createAdminMenuRequest)
	res := new(createAdminMenuResponse)
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

	createData := new(admin.CreateMenuData)
	createData.AdminId = int64(ids[0])
	createData.Actions = req.Actions

	err = h.adminService.CreateMenu(ctx, createData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminMenuCreateError, err)
		return
	}

	res.Id = int64(ids[0])
	api.ResponseOK(ctx, res)
}

package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_login.html", nil)
}

func (h *handler) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_add.html", nil)
}

func (h *handler) List(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_list.html", nil)
}

func (h *handler) Menu(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "menu_view.html", nil)
}

func (h *handler) AdminMenu(ctx *gin.Context) {
	type adminMenuRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type adminMenuResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	req := new(adminMenuRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	obj := new(adminMenuResponse)
	obj.HashID = req.Id

	ctx.HTML(http.StatusOK, "admin_menu.html", obj)
}

func (h *handler) MenuAction(ctx *gin.Context) {
	type menuActionRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type menuActionResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	req := new(menuActionRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	obj := new(menuActionResponse)
	obj.HashID = req.Id

	ctx.HTML(http.StatusOK, "menu_action.html", obj)
}

func (h *handler) ModifyInfo(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_modify_info.html", nil)
}

func (h *handler) ModifyPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_modify_password.html", nil)
}

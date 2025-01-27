package menu

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type createRequest struct {
	Id    string `form:"id"`    // ID
	Pid   int64  `form:"pid"`   // 父类ID
	Name  string `form:"name"`  // 菜单名称
	Link  string `form:"link"`  // 链接地址
	Icon  string `form:"icon"`  // 图标
	Level int32  `form:"level"` // 菜单类型 1:一级菜单 2:二级菜单
}

type createResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Create 创建/编辑菜单
// @Summary 创建/编辑菜单
// @Description 创建/编辑菜单
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu [post]
// @Security LoginToken
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetError(err).Error())
		return
	}

	if req.Id != "" { // 编辑功能
		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
			return
		}

		id := int64(ids[0])

		updateData := new(menu.UpdateMenuData)
		updateData.Name = req.Name
		updateData.Icon = req.Icon
		updateData.Link = req.Link

		err = h.service.Modify(ctx, id, updateData)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.MenuUpdateError, err)
			return
		}

		res.Id = id
		api.ResponseOK(ctx, res)

	} else { // 新增功能

		pid := int64(req.Level)
		level := 2

		if req.Level == -1 {
			pid = 0
			level = 1
		}

		createData := new(menu.CreateMenuData)
		createData.Pid = pid
		createData.Name = req.Name
		createData.Icon = req.Icon
		createData.Link = req.Link
		createData.Level = cast.ToInt32(level)

		id, err := h.service.Create(ctx, createData)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.MenuCreateError, err)
			return
		}

		res.Id = id
		api.ResponseOK(ctx, res)
	}
}

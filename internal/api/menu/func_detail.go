package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Id   int64  `json:"id"`   // 主键ID
	Pid  int64  `json:"pid"`  // 父类ID
	Name string `json:"name"` // 菜单名称
	Link string `json:"link"` // 链接地址
	Icon string `json:"icon"` // 图标
}

// Detail 菜单详情
// @Summary 菜单详情
// @Description 菜单详情
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/{id} [get]
// @Security LoginToken
func (h *handler) Detail(ctx *gin.Context) {
	req := new(detailRequest)
	res := new(detailResponse)
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

	id := int64(ids[0])

	searchOneData := new(menu.SearchOneData)
	searchOneData.Id = id

	info, err := h.service.Detail(ctx, searchOneData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuDetailError, err)
		return
	}

	res.Id = info.Id
	res.Pid = info.Pid
	res.Name = info.Name
	res.Link = info.Link
	res.Icon = info.Icon
	api.ResponseOK(ctx, res)
}

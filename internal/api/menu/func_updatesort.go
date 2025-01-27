package menu

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updateSortRequest struct {
	Id   string `form:"id"`   // HashId
	Sort int32  `form:"sort"` // 排序
}

type updateSortResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// UpdateSort 更新菜单排序
// @Summary 更新菜单排序
// @Description 更新菜单排序
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashId"
// @Param sort formData int true "排序"
// @Success 200 {object} updateSortResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/sort [patch]
// @Security LoginToken
func (h *handler) UpdateSort(ctx *gin.Context) {
	req := new(updateSortRequest)
	res := new(updateSortResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetError(err).Error())
		return
	}

	ids, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	id := int64(ids[0])

	err = h.service.UpdateSort(ctx, id, req.Sort)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuUpdateError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

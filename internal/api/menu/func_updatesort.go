package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateSortRequest struct {
	Id   string `form:"id"`   // HashId
	Sort int32  `form:"sort"` // 排序
}

type updateSortResponse struct {
	Id int32 `json:"id"` // 主键ID
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

	err = h.menuService.UpdateSort(ctx, id, req.Sort)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuUpdateError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

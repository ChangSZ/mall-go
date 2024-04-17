package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Tags API.menu
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/{id} [delete]
// @Security LoginToken
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	res := new(deleteResponse)
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

	err = h.menuService.Delete(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuDeleteError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

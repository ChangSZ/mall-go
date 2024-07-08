package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type deleteActionRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteActionResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// DeleteAction 删除功能权限
// @Summary 删除功能权限
// @Description 删除功能权限
// @Tags API.menu
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteActionResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu_action/{id} [delete]
// @Security LoginToken
func (h *handler) DeleteAction(ctx *gin.Context) {
	req := new(deleteActionRequest)
	res := new(deleteActionResponse)
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

	err = h.service.DeleteAction(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuDeleteActionError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

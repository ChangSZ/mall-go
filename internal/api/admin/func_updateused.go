package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateUsedRequest struct {
	Id   string `form:"id"`   // 主键ID
	Used int32  `form:"used"` // 是否启用 1:是 -1:否
}

type updateUsedResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// UpdateUsed 更新管理员为启用/禁用
// @Summary 更新管理员为启用/禁用
// @Description 更新管理员为启用/禁用
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "Hashid"
// @Param used formData int true "是否启用 1:是 -1:否"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/used [patch]
// @Security LoginToken
func (h *handler) UpdateUsed(ctx *gin.Context) {
	req := new(updateUsedRequest)
	res := new(updateUsedResponse)
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

	id := int64(ids[0])
	err = h.service.UpdateUsed(ctx, id, req.Used)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminUpdateError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

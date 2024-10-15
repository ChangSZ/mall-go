package admin

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type resetPasswordRequest struct {
	Id string `uri:"id"` // HashID
}

type resetPasswordResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// ResetPassword 重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags API.admin
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} resetPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/reset_password/{id} [patch]
// @Security LoginToken
func (h *handler) ResetPassword(ctx *gin.Context) {
	req := new(resetPasswordRequest)
	res := new(resetPasswordResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
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
	err = h.service.ResetPassword(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminResetPasswordError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

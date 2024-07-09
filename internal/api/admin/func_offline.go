package admin

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type offlineRequest struct {
	Id string `form:"id"` // 主键ID
}

type offlineResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Offline 下线管理员
// @Summary 下线管理员
// @Description 下线管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "Hashid"
// @Success 200 {object} offlineResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/offline [patch]
// @Security LoginToken
func (h *handler) Offline(ctx *gin.Context) {
	req := new(offlineRequest)
	res := new(offlineResponse)
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

	b := redis.Cache().Del(ctx, configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id))
	if !b {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminOfflineError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)

}

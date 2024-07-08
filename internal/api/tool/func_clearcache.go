package tool

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type clearCacheRequest struct {
	RedisKey string `form:"redis_key"` // Redis Key
}

type clearCacheResponse struct {
	Bool bool `json:"bool"` // 删除结果
}

// ClearCache 清空缓存
// @Summary 清空缓存
// @Description 清空缓存
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param redis_key formData string true "Redis Key"
// @Success 200 {object} searchCacheResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/cache/clear [patch]
// @Security LoginToken
func (h *handler) ClearCache(ctx *gin.Context) {
	req := new(clearCacheRequest)
	res := new(clearCacheResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	if b := redis.Cache().Exists(ctx, req.RedisKey); !b {
		api.Response(ctx, http.StatusBadRequest, code.CacheNotExist)
		return
	}

	if b := redis.Cache().Del(ctx, req.RedisKey); !b {
		api.Response(ctx, http.StatusBadRequest, code.CacheDelError)
		return
	}

	res.Bool = true
	api.ResponseOK(ctx, res)
}

package tool

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type searchCacheRequest struct {
	RedisKey string `form:"redis_key"` // Redis Key
}

type searchCacheResponse struct {
	Val string `json:"val"` // 查询后的值
	TTL string `json:"ttl"` // 过期时间
}

// SearchCache 查询缓存
// @Summary 查询缓存
// @Description 查询缓存
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param redis_key formData string true "Redis Key"
// @Success 200 {object} searchCacheResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/cache/search [post]
// @Security LoginToken
func (h *handler) SearchCache(ctx *gin.Context) {
	req := new(searchCacheRequest)
	res := new(searchCacheResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	if b := redis.Cache().Exists(ctx, req.RedisKey); !b {
		api.Response(ctx, http.StatusBadRequest, code.CacheNotExist)
		return
	}

	val, err := redis.Cache().Get(ctx, req.RedisKey)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.CacheGetError, err)
		return
	}

	ttl, err := redis.Cache().TTL(ctx, req.RedisKey)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.CacheGetError, err)
		return
	}

	res.Val = val
	res.TTL = ttl.String()
	api.ResponseOK(ctx, res)
}

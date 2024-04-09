package tool

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
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
func (h *handler) SearchCache() core.HandlerFunc {
	return func(c core.Context) {
		req := new(searchCacheRequest)
		res := new(searchCacheResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		if b := redis.Cache().Exists(req.RedisKey); !b {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CacheNotExist,
				code.Text(code.CacheNotExist)),
			)
			return
		}

		val, err := redis.Cache().Get(req.RedisKey, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithError(err),
			)
			return
		}

		ttl, err := redis.Cache().TTL(req.RedisKey)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithError(err),
			)
			return
		}

		res.Val = val
		res.TTL = ttl.String()
		c.Payload(res)
	}
}

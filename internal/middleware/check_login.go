package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
)

// CheckLogin 验证是否登录
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(configs.HeaderLoginToken)
		if token == "" {
			err := errors.New("header 中缺少 Token 参数")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		if !redis.Cache().Exists(ctx, configs.RedisKeyPrefixLoginUser+token) {
			err := errors.New("请先登录")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		cacheData, cacheErr := redis.Cache().Get(ctx, configs.RedisKeyPrefixLoginUser+token)
		if cacheErr != nil {
			err := errors.New("header 中缺少 Token 参数")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.AuthorizationError, cacheErr)
			ctx.Abort()
			return
		}

		var sessionUserInfo = proposal.SessionUserInfo{}
		jsonErr := json.Unmarshal([]byte(cacheData), &sessionUserInfo)
		if jsonErr != nil {
			log.WithTrace(ctx).Error(jsonErr)
			api.Response(ctx, http.StatusUnauthorized, code.AuthorizationError, jsonErr)
			ctx.Abort()
			return
		}

		core.SetSessionUserInfo(ctx, sessionUserInfo)
		ctx.Next()
	}
}

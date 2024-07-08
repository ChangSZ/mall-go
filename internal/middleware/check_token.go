package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/ums_admin"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"
	"github.com/ChangSZ/mall-go/pkg/jwt"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

var (
	jwtConfig    = configs.Get().Jwt
	jwtTokenUtil = jwt.NewJwtTokenUtil(jwtConfig.Secret, jwtConfig.Expiration, jwtConfig.TokenHead)
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(jwtConfig.TokenHeader)
		if token == "" {
			err := errors.New("header中缺少token参数")
			log.WithTrace(ctx).Error(err)
			api.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}
		token = strings.TrimPrefix(token, jwtConfig.TokenHead)
		username := jwtTokenUtil.GetUserNameFromToken(token)
		if username == "" {
			api.Unauthorized(ctx, "token非法, 无用户信息")
			ctx.Abort()
			return
		}
		userDetails, loadErr := ums_admin.New().LoadUserByUsername(ctx, username)
		if loadErr != nil {
			err := fmt.Errorf("未找到用户: %v, %w", username, loadErr)
			log.WithTrace(ctx).Error(err)
			api.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}
		if validateErr := jwtTokenUtil.ValidateToken(token, userDetails.GetUsername()); validateErr != nil {
			log.WithTrace(ctx).Error(validateErr)
			api.Unauthorized(ctx, validateErr.Error())
			ctx.Abort()
			return
		}
		var userInfo = proposal.UmsUserInfo{
			UserName: userDetails.GetUsername(),
			Token:    token,
		}
		core.SetUmsUserInfo(ctx, userInfo)
		ctx.Next()
	}
}

func CheckMemberToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(jwtConfig.TokenHeader)
		if token == "" {
			err := errors.New("header中缺少token参数")
			log.WithTrace(ctx).Error(err)
			api.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}
		token = strings.TrimPrefix(token, jwtConfig.TokenHead)
		username := jwtTokenUtil.GetUserNameFromToken(token)
		if username == "" {
			api.Unauthorized(ctx, "token非法, 无用户信息")
			ctx.Abort()
			return
		}
		userDetails, loadErr := ums_member.New().LoadUserByUsername(ctx, username)
		if loadErr != nil {
			err := fmt.Errorf("未找到用户: %v, %w", username, loadErr)
			log.WithTrace(ctx).Error(err)
			api.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}
		if validateErr := jwtTokenUtil.ValidateToken(token, userDetails.GetUsername()); validateErr != nil {
			log.WithTrace(ctx).Error(validateErr)
			api.Unauthorized(ctx, validateErr.Error())
			ctx.Abort()
			return
		}
		var userInfo = proposal.UmsUserInfo{
			UserName: userDetails.GetUsername(),
			Token:    token,
		}
		core.SetUmsUserInfo(ctx, userInfo)
		ctx.Next()
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized"
	authorizedService "github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/ChangSZ/mall-go/pkg/env"
	"github.com/ChangSZ/mall-go/pkg/errors"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/signature"
	"github.com/ChangSZ/mall-go/pkg/urltable"
	"github.com/gin-gonic/gin"
)

var whiteListPath = map[string]bool{
	"/login/web": true,
}

// CheckSignature 验证签名是否合法，对用签名算法 pkg/signature
func CheckSignature() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !env.Active().IsPro() {
			ctx.Next()
			return
		}

		// 签名信息
		authorization := ctx.GetHeader(configs.HeaderSignToken)
		if authorization == "" {
			err := errors.New("Header 中缺少 Authorization 参数")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		// 时间信息
		date := ctx.GetHeader(configs.HeaderSignTokenDate)
		if date == "" {
			err := errors.New("Header 中缺少 Date 参数")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		// 通过签名信息获取 key
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < 2 {
			err := errors.New("Header 中 Authorization 格式错误")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		key := authorizationSplit[0]

		data, err := authorizedService.New().DetailByKey(ctx, key)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		if data.IsUsed == authorized.IsUsedNo {
			err := errors.New(key + " 已被禁止调用")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		if len(data.Apis) < 1 {
			err := errors.New(key + " 未进行接口授权")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		urlPath := ctx.Request.URL.Path
		method := ctx.Request.Method
		if !whiteListPath[urlPath] {
			// 验证 method + urlPath 是否授权
			table := urltable.NewTable()
			for _, v := range data.Apis {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(method + urlPath); pattern == "" {
				err := errors.New(method + urlPath + " 未进行接口授权")
				log.WithTrace(ctx).Error(err)
				api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
				ctx.Abort()
				return
			}
		}
		_ = ctx.Request.ParseForm()
		ok, err := signature.New(key, data.Secret, configs.HeaderSignTokenTimeout).Verify(authorization, date, urlPath, method, ctx.Request.Form)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		if !ok {
			err := errors.New("Header 中 Authorization 信息错误")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.AuthorizationError, err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

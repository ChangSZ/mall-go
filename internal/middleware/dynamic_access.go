package middleware

import (
	"net/http"
	"path/filepath"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/ums_admin"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

func match(pattern, path string) (bool, error) {
	return filepath.Match(pattern, path)
}

func DynamicAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userInfo := core.GetUmsUserInfo(ctx)
		userDetails, err := ums_admin.New().LoadUserByUsername(ctx, userInfo.UserName)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Failed(ctx, err.Error())
			ctx.Abort()
		}
		url := ctx.Request.URL.Path
		for _, v := range userDetails.ResourceList {
			matched, _ := match(v.Url, url)
			if matched {
				ctx.Next()
				return
			}
		}
		api.Response(ctx, http.StatusForbidden, code.FORBIDDEN, "抱歉，您没有访问权限")
		ctx.Abort()
	}
}

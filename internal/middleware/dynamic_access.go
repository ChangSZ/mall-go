package middleware

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/ums_admin"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/glob"
)

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
			g := glob.MustCompile(v.Url)
			if g.Match(url) {
				ctx.Next()
				return
			}
		}
		api.Response(ctx, http.StatusForbidden, code.FORBIDDEN, "抱歉，您没有访问权限")
		ctx.Abort()
	}
}

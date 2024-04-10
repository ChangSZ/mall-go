package middleware

import (
	"context"

	"github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware"
)

// AddTraceCtx: kratos的中间件, 用来将trace的context添加到gin.Context中
func AddTraceCtx(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if c, ok := gin.FromGinContext(ctx); ok {
			c.Request = c.Request.WithContext(ctx)
		}
		reply, err = handler(ctx, req)
		return
	}
}

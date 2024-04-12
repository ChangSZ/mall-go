package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/alert"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/pkg/env"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

func AlertNotify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}
		defer func() {
			// region 发生 Panic 异常发送告警提醒
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				log.Errorf("got panic: %+v, stack: %v", err, stackInfo)
				core.AbortWithError(ctx, core.Error(
					http.StatusInternalServerError,
					code.ServerError,
					code.Text(code.ServerError)))

				alert.NotifyHandler()(&proposal.AlertMessage{
					ProjectName:  configs.ProjectName,
					Env:          env.Active().Value(),
					TraceID:      core.TraceID(ctx),
					HOST:         ctx.Request.Host,
					URI:          ctx.Request.URL.Path,
					Method:       ctx.Request.Method,
					ErrorMessage: err,
					ErrorStack:   stackInfo,
					Timestamp:    time.Now(),
				})
			}
			// endregion

			// region 发生错误，进行返回
			// if ctx.IsAborted() {
			// 	if err := core.AbortError(ctx); err != nil { // customer err
			// 		// 判断是否需要发送告警通知
			// 		if err.IsAlert() {
			// 			alert.NotifyHandler()(&proposal.AlertMessage{
			// 				ProjectName:  configs.ProjectName,
			// 				Env:          env.Active().Value(),
			// 				TraceID:      core.TraceID(ctx),
			// 				HOST:         ctx.Request.Host,
			// 				URI:          ctx.Request.URL.Path,
			// 				Method:       ctx.Request.Method,
			// 				ErrorMessage: err.Message(),
			// 				ErrorStack:   fmt.Sprintf("%+v", err.StackError()),
			// 				Timestamp:    time.Now(),
			// 			})
			// 		}

			// 		response := &code.Failure{
			// 			Code:    err.BusinessCode(),
			// 			Message: err.Message(),
			// 		}
			// 		ctx.JSON(err.HTTPCode(), response)
			// 	}
			// }
			// endregion
		}()

		ctx.Next()
	}
}

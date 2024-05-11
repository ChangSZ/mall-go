package core

import (
	"github.com/gin-gonic/gin"
)

// AliasForRecordMetrics 对请求路径起个别名，用于记录指标。
// 如：Get /user/:username 这样的路径，因为 username 会有非常多的情况，这样记录指标非常不友好。
func AliasForRecordMetrics(path string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		SetAlias(ctx, path)
	}
}

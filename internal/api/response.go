package api

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpCode, errCode int, data ...interface{}) {
	if len(data) > 0 {
		ctx.JSON(httpCode, gin.H{
			"code":    errCode,
			"message": code.Text(errCode),
			"data":    data[0],
		})
		return
	}

	ctx.JSON(httpCode, gin.H{
		"code":    errCode,
		"message": code.Text(errCode),
	})
}

func ResponseOK(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"data": data[0],
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

package api

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/code"

	"github.com/gin-gonic/gin"
)

// 框架失败或者mall请使用这个函数
func Response(ctx *gin.Context, httpCode, errCode int, data ...interface{}) {
	res := code.Failure{
		Code:    errCode,
		Message: code.Text(errCode),
	}
	if len(data) > 0 {
		res.Data = data[0]
		ctx.JSON(httpCode, res)
		return
	}
	ctx.JSON(httpCode, res)
}

// 这个函数是给框架来使用的
func ResponseOK(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 {
		ctx.JSON(http.StatusOK, data[0])
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func Success(ctx *gin.Context, data ...interface{}) {
	res := code.Success{
		Code:    http.StatusOK,
		Message: "操作成功",
		Data:    nil,
	}
	if len(data) > 0 {
		res.Data = data[0]
		ctx.JSON(http.StatusOK, res)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/code"
)

// 框架失败请使用这个函数
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

func response(ctx *gin.Context, codeNum int, data interface{}) {
	res := code.Success{
		Code:    codeNum,
		Message: code.Text(codeNum),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, res)
}

func getData(data []interface{}) interface{} {
	if len(data) > 0 {
		return data[0]
	}
	return nil
}

func Success(ctx *gin.Context, data ...interface{}) {
	response(ctx, code.SUCCESS, getData(data))
}

func Failed(ctx *gin.Context, data ...interface{}) {
	response(ctx, code.FAILED, getData(data))
}

func ValidateFailed(ctx *gin.Context, data ...interface{}) {
	response(ctx, code.VALIDATE_FAILED, getData(data))
}

func Unauthorized(ctx *gin.Context, data ...interface{}) {
	response(ctx, code.UNAUTHORIZED, getData(data))
}

func Forbidden(ctx *gin.Context, data ...interface{}) {
	response(ctx, code.FORBIDDEN, getData(data))
}

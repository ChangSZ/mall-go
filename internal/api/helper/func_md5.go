package helper

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/gin-gonic/gin"
)

type md5Request struct {
	Str string `uri:"str" binding:"required"` // 需要加密的字符串
}

type md5Response struct {
	Md5Str string `json:"md5_str"` // MD5后的字符串
}

// Md5 加密
// @Summary 加密
// @Description 加密
// @Tags Helper
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param str path string true "需要加密的字符串"
// @Success 200 {object} md5Response
// @Failure 400 {object} code.Failure
// @Router /helper/md5/{str} [get]
func (h *handler) Md5(ctx *gin.Context) {
	req := new(md5Request)
	res := new(md5Response)

	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	m := md5.New()
	m.Write([]byte(req.Str))
	res.Md5Str = hex.EncodeToString(m.Sum(nil))
	api.ResponseOK(ctx, res)
}

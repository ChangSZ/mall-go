package ums_admin

import (
	"net/http"
	"time"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/ums_admin"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Icon     string `json:"icon"`
	Email    string `json:"email" binding:"email"`
	NickName string `json:"nickName"`
	Note     string `json:"note"`
}

type registerResponse struct {
	Id         int64     //
	Username   string    //
	Password   string    //
	Icon       string    // 头像
	Email      string    // 邮箱
	NickName   string    // 昵称
	Note       string    // 备注信息
	CreateTime time.Time // 创建时间
	LoginTime  time.Time // 最后登录时间
	Status     int32     // 帐号启用状态：0->禁用；1->启用
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body registerRequest true "请求信息"
// @Success 200 {object} code.Success{data=registerResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/register [post]
func (h *handler) Register(ctx *gin.Context) {
	req := new(registerRequest)
	res := new(registerResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	umsAdminParam := new(ums_admin.UmsAdminParam)
	umsAdminParam.Username = req.Username
	umsAdminParam.Password = req.Password
	umsAdminParam.Icon = req.Icon
	umsAdminParam.Email = req.Email
	umsAdminParam.NickName = req.NickName
	umsAdminParam.Note = req.Note

	umsAdmin, err := h.umsAdminService.Register(ctx, umsAdminParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminRegisterError, err)
		return
	}
	res.Id = umsAdmin.Id
	res.Username = umsAdmin.Username
	res.Password = umsAdmin.Password
	res.Icon = umsAdmin.Icon
	res.Email = umsAdmin.Email
	res.NickName = umsAdmin.NickName
	res.Note = umsAdmin.Note
	res.CreateTime = umsAdmin.CreateTime
	res.LoginTime = umsAdmin.LoginTime
	res.Status = umsAdmin.Status
	api.Success(ctx, res)
}

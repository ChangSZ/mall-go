package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Username string `form:"username" binding:"required"` // 用户名
	Nickname string `form:"nickname" binding:"required"` // 昵称
	Mobile   string `form:"mobile" binding:"required"`   // 手机号
	Password string `form:"password" binding:"required"` // MD5后的密码
}

type createResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Create 新增管理员
// @Summary 新增管理员
// @Description 新增管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param nickname formData string true "昵称"
// @Param mobile formData string true "手机号"
// @Param password formData string true "MD5后的密码"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
// @Security LoginToken
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	createData := new(admin.CreateAdminData)
	createData.Nickname = req.Nickname
	createData.Username = req.Username
	createData.Mobile = req.Mobile
	createData.Password = req.Password

	id, err := h.adminService.Create(ctx, createData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminCreateError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

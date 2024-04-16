package ums_admin

import (
	"net/http"
	"time"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Keyword  string `json:"keyword" binding:"omitempty"`
	PageSize int    `json:"pageSize" binding:"omitempty"`
	PageNum  int    `json:"pageNum" binding:"omitempty"`
}

type UmsAdmin struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Icon       string    `json:"icon"`
	Email      string    `json:"email"`
	NickName   string    `json:"nickName"`
	Note       string    `json:"note"`
	CreateTime time.Time `json:"createTime"`
	LoginTime  time.Time `json:"loginTime"`
	Status     int32     `json:"status"`
}

type listResponse struct {
	PageNum   int        `json:"pageNum"`
	PageSize  int        `json:"pageSize"`
	TotalPage int64      `json:"totalPage"`
	Total     int64      `json:"total"`
	List      []UmsAdmin `json:"list"`
}

// List 根据用户名或姓名分页获取用户列表
// @Summary 根据用户名或姓名分页获取用户列表
// @Description 根据用户名或姓名分页获取用户列表
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}
	if req.PageSize == 0 {
		req.PageSize = 5
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	list, total, err := h.umsAdminService.List(ctx, req.Keyword, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminGetListError, err)
		return
	}
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	totalPage := total / int64(req.PageSize)
	if total%int64(req.PageSize) > 0 {
		totalPage += 1
	}
	res.TotalPage = totalPage
	res.Total = total
	listData := make([]UmsAdmin, 0, len(list))
	for _, v := range list {
		listData = append(listData, UmsAdmin{
			ID:         v.Id,
			Username:   v.Username,
			Password:   v.Password,
			Icon:       v.Icon,
			Email:      v.Email,
			NickName:   v.NickName,
			Note:       v.Note,
			CreateTime: v.CreateTime,
			LoginTime:  v.LoginTime,
			Status:     v.Status,
		})
	}
	res.List = listData
	api.Success(ctx, res)
}

package cron

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/cron"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type modifyRequest struct {
	Id                  string `form:"id" binding:"required"`             // 任务ID
	Name                string `form:"name" binding:"required"`           // 任务名称
	Spec                string `form:"spec" binding:"required"`           // crontab 表达式
	Command             string `form:"command" binding:"required"`        // 执行命令
	Protocol            int32  `form:"protocol" binding:"required"`       // 执行方式 1:shell 2:http
	HttpMethod          int32  `form:"http_method"`                       // http 请求方式 1:get 2:post
	Timeout             int32  `form:"timeout" binding:"required"`        // 超时时间(单位:秒)
	RetryTimes          int32  `form:"retry_times" binding:"required"`    // 重试次数
	RetryInterval       int32  `form:"retry_interval" binding:"required"` // 重试间隔(单位:秒)
	NotifyStatus        int32  `form:"notify_status" binding:"required"`  // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyType          int32  `form:"notify_type"`                       // 通知类型 1:邮件 2:webhook
	NotifyReceiverEmail string `form:"notify_receiver_email"`             // 通知者邮箱地址(多个用,分割)
	NotifyKeyword       string `form:"notify_keyword"`                    // 通知匹配关键字(多个用,分割)
	Remark              string `form:"remark"`                            // 备注
	IsUsed              int32  `form:"is_used" binding:"required"`        // 是否启用 1:是  -1:否
}

type modifyResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Modify 编辑任务
// @Summary 编辑任务
// @Description 编辑任务
// @Tags API.cron
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashID"
// @Param name formData string true "任务名称"
// @Param spec formData string true "crontab 表达式"
// @Param command formData string true "执行命令"
// @Param protocol formData int true "执行方式 1:shell 2:http"
// @Param http_method formData int false "http 请求方式 1:get 2:post"
// @Param timeout formData int true "超时时间(单位:秒)"
// @Param retry_times formData int true "重试次数"
// @Param retry_interval formData int true "重试间隔(单位:秒)"
// @Param notify_status formData int true "执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知"
// @Param notify_type formData int false "通知类型 1:邮件 2:webhook"
// @Param notify_receiver_email formData string false "通知者邮箱地址(多个用,分割)"
// @Param notify_keyword formData string false "通知匹配关键字(多个用,分割)"
// @Param remark formData string false "备注"
// @Param is_used formData int true "是否启用 1:是  -1:否"
// @Success 200 {object} modifyResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/{id} [post]
// @Security LoginToken
func (h *handler) Modify(ctx *gin.Context) {
	req := new(modifyRequest)
	res := new(modifyResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetError(err).Error())
		return
	}

	ids, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	id := int64(ids[0])

	modifyData := new(cron.ModifyCronTaskData)
	modifyData.Name = req.Name
	modifyData.Spec = req.Spec
	modifyData.Command = req.Command
	modifyData.Protocol = req.Protocol
	modifyData.HttpMethod = req.HttpMethod
	modifyData.Timeout = req.Timeout
	modifyData.RetryTimes = req.RetryTimes
	modifyData.RetryInterval = req.RetryInterval
	modifyData.NotifyStatus = req.NotifyStatus
	modifyData.NotifyType = req.NotifyType
	modifyData.NotifyReceiverEmail = req.NotifyReceiverEmail
	modifyData.NotifyKeyword = req.NotifyKeyword
	modifyData.Remark = req.Remark
	modifyData.IsUsed = req.IsUsed

	if err := h.service.Modify(ctx, id, modifyData); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.CronUpdateError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}

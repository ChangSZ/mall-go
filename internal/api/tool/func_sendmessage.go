package tool

import (
	"encoding/json"
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/timeutil"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/websocket/sysmessage"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type sendMessageRequest struct {
	Message string `form:"message"` // 消息内容
}

type sendMessageResponse struct {
	Status string `json:"status"` // 状态
}

// SendMessage 发送消息
// @Summary 发送消息
// @Description 发送消息
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param message formData string true "消息内容"
// @Success 200 {object} sendMessageResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/send_message [post]
// @Security LoginToken
func (h *handler) SendMessage(ctx *gin.Context) {
	type messageBody struct {
		Username string `json:"username"`
		Message  string `json:"message"`
		Time     string `json:"time"`
	}
	req := new(sendMessageRequest)
	res := new(sendMessageResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetError(err).Error())
		return
	}

	conn, err := sysmessage.GetConn()
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.SocketConnectError, err)
		return
	}

	messageData := new(messageBody)
	messageData.Username = core.SessionUserInfo(ctx).UserName
	messageData.Message = req.Message
	messageData.Time = timeutil.CSTLayoutString()

	messageJsonData, err := json.Marshal(messageData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.SocketSendError, err)
		return
	}

	err = conn.OnSend(messageJsonData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.SocketSendError, err)
		return
	}

	res.Status = "OK"
	api.ResponseOK(ctx, res)
}

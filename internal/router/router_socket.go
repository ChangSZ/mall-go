package router

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/websocket/sysmessage"
)

func setSocketRouter(r *resource) {
	systemMessage := sysmessage.New(r.logger)

	// 无需记录日志
	socket := r.mux.Group("/socket", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 系统消息
		socket.GET("/system/message", systemMessage.Connect())
	}
}

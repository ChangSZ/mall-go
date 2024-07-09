package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/websocket/sysmessage"
)

func setSocketRouter(eng *gin.Engine) {
	systemMessage := sysmessage.New()

	// 无需记录日志
	socket := eng.Group("/socket")
	{
		// 系统消息
		socket.GET("/system/message", systemMessage.Connect)
	}
}

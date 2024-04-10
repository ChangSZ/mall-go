package router

import (
	"github.com/ChangSZ/mall-go/internal/websocket/sysmessage"
	"github.com/gin-gonic/gin"
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

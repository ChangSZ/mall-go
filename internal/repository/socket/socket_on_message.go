package socket

import (
	"github.com/ChangSZ/mall-go/pkg/log"
)

func (s *server) OnMessage() {
	defer func() {
		s.OnClose()
	}()

	for {
		//接收消息
		_, message, err := s.socket.ReadMessage()
		if err != nil {
			log.Error("socket on message error: ", err)
			break
		}

		// 为了便于演示，仅输出到日志文件
		log.Info("receive message: " + string(message))
	}
}

package socket

import (
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gorilla/websocket"
)

func (s *server) OnSend(message []byte) error {
	err := s.socket.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		s.OnClose()
		log.Error("socket on send error: ", err)
	}
	return err
}

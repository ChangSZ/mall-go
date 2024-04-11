package socket

import "github.com/ChangSZ/mall-go/pkg/log"

func (s *server) OnClose() {
	err := s.socket.Close()
	if err != nil {
		log.Error("socket on closed error: ", err)
	}
}

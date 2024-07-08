package socket

import "github.com/ChangSZ/golib/log"

func (s *server) OnClose() {
	err := s.socket.Close()
	if err != nil {
		log.Error("socket on closed error: ", err)
	}
}

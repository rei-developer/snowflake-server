package server

import (
	"net"

	"github.com/snowflake-server/src/handlers"
)

const (
	loginVerification MessageType = iota
	drawFirstLover
	getUserByID
)

func (s *Server) processIncomingMessages(conn net.Conn, msg Message, outgoing chan []byte) {
	//s.mu.Lock()
	//defer s.mu.Unlock()

	println("왔니")
	println(msg.Type)

	switch msg.Type {
	case loginVerification:
		println("왜 왔니")
		if handlers.HandleLoginVerification(conn, msg.Payload, s.users, &s.nextUserIndex) {
			s.NotifyUserList()
		}
	case drawFirstLover:
		println("왔니 여기도")

		handlers.HandleDrawFirstLover(msg.Payload, outgoing)
	case getUserByID:
		handlers.HandleGetUserByID(msg.Payload, outgoing, s.users)
	default:
		handlers.HandleUnknownPacket(uint32(msg.Type))
	}
}

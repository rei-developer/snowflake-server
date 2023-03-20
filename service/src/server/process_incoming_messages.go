package server

import (
	"net"

	"github.com/snowflake-server/src/handlers"
)

const (
	loginVerification MessageType = iota + 1
	drawFirstLover
	getUserByID
)

func (s *Server) processIncomingMessages(conn net.Conn, msg Message, outgoing chan []byte) {
	//s.mu.Lock()
	//defer s.mu.Unlock()

	switch msg.Type {
	case loginVerification:
		if handlers.HandleLoginVerification(conn, msg.Payload, s.users, &s.nextUserIndex) {
			s.NotifyUserList()
		}
	case drawFirstLover:
		handlers.HandleDrawFirstLover(msg.Payload, outgoing)
	case getUserByID:
		handlers.HandleGetUserByID(msg.Payload, outgoing, s.users)
	default:
		if msg.Type == 0 {
			break
		}
		handlers.HandleUnknownPacket(uint32(msg.Type))
	}
}

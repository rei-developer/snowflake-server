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
	switch msg.Type {
	case loginVerification:
		if handlers.HandleLoginVerification(conn, msg.Payload, s.users, &s.nextUserIndex) {
			s.NotifyUserList()
		}
	case drawFirstLover:
		handlers.HandleDrawFirstLover(conn, msg.Payload)
	case getUserByID:
		handlers.HandleGetUserByID(msg.Payload, s.users, outgoing, &s.mu)
	default:
		handlers.HandleUnknownPacket(uint32(msg.Type))
	}
}

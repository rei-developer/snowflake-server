package server

import (
	"net"

	"github.com/snowflake-server/src/handlers"
)

func (s *Server) processIncomingMessages(conn net.Conn, incoming chan Message, outgoing chan []byte) {
	for msg := range incoming {
		switch msg.Type {
		case loginVerification:
			handlers.HandleLoginVerification(conn, msg.Payload, s.users, &s.nextUserIndex, &s.mu)
		case getUserByID:
			handlers.HandleGetUserByID(msg.Payload, s.users, outgoing, &s.mu)
		case hello:
			handlers.HandleHello(outgoing)
		default:
			handlers.HandleUnknownPacket(uint32(msg.Type))
		}
	}
}

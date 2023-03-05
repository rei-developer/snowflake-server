package server

import (
	"net"

	"github.com/snowflake-server/src/handlers"
)

func (s *Server) processIncomingMessages(conn net.Conn, incoming chan Message, outgoing chan []byte) {
	for msg := range incoming {
		switch msg.Type {
		case getUserByID:
			handlers.HandleGetUserByID(msg.Payload, &s.mu, s.users, outgoing)
		case hello:
			handlers.HandleHello(outgoing)
		default:
			handlers.HandleUnknownPacket(uint32(msg.Type))
		}
	}
}

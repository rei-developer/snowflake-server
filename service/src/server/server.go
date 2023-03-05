package server

import (
	"fmt"
	"github.com/snowflake-server/src/handlers"
	"net"
	"sync"

	"github.com/snowflake-server/src/user"
)

type Server struct {
	listener      net.Listener
	users         map[uint32]*user.User
	nextUserIndex uint32
	mu            sync.Mutex
}

func NewServer(port string) (*Server, error) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener:      l,
		users:         make(map[uint32]*user.User),
		nextUserIndex: 1,
	}, nil
}

func (s *Server) Start() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) Stop() error {
	return s.listener.Close()
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
		handlers.HandleUserDisconnect(conn, s.users, &s.mu)
	}(conn)

	incoming := make(chan Message)
	outgoing := make(chan []byte, 10)

	go s.readIncomingMessages(conn, incoming, outgoing)
	go s.writeOutgoingMessages(conn, outgoing)
	s.processIncomingMessages(conn, incoming, outgoing)
}

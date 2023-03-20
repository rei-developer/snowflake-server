package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/snowflake-server/src/handlers"
	"github.com/snowflake-server/src/user"
)

type Server struct {
	listener      net.Listener
	users         map[uint32]*user.User
	nextUserIndex uint32
	mu            sync.Mutex
	userObservers []UserObserver
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
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Println("Error converting conn to TCPConn")
		return
	}
	if err := tcpConn.SetKeepAlive(true); err != nil {
		fmt.Printf("Error enabling keep-alive: %v\n", err)
		return
	}
	if err := tcpConn.SetKeepAlivePeriod(5 * time.Minute); err != nil {
		fmt.Printf("Error setting keep-alive period: %v\n", err)
		return
	}

	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			fmt.Printf("Error closing connection: %v\n", err)
		}
		handlers.HandleUserDisconnect(conn, s.users, &s.mu)
	}(conn)

	incoming := make(chan Message)
	outgoing := make(chan []byte, 10)

	go s.readIncomingMessages(conn, incoming, outgoing)

	for {
		select {
		case msg := <-incoming:
			s.processIncomingMessages(conn, msg, outgoing)
		case data := <-outgoing:
			_, err := conn.Write(data)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Connection closed by client")
				} else {
					fmt.Printf("Error writing data: %v\n", err)
				}
				return
			}
		}
	}
}

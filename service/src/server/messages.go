package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	hello MessageType = iota
	getUserByID
)

type MessageType uint32

type Message struct {
	Type    MessageType
	Payload []byte
}

func (s *Server) readIncomingMessages(conn net.Conn, incoming chan<- Message, outgoing chan<- []byte) {
	defer close(incoming)
	defer close(outgoing)

	for {
		header := make([]byte, 8)
		_, err := io.ReadFull(conn, header)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Println("Error reading header:", err)
			}
			return
		}

		messageType := MessageType(binary.BigEndian.Uint32(header[:4]))
		payloadLength := binary.BigEndian.Uint32(header[4:])

		payload := make([]byte, payloadLength)
		_, err = io.ReadFull(conn, payload)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Println("Error reading payload:", err)
			}
			return
		}

		incoming <- Message{Type: messageType, Payload: payload}
	}
}

func (s *Server) writeOutgoingMessages(conn net.Conn, outgoing <-chan []byte) {
	for data := range outgoing {
		_, err := conn.Write(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Println("Error writing data:", err)
			}
			return
		}
	}
}

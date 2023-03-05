package handlers

import (
	"encoding/json"
	"github.com/snowflake-server/src/user"
	"net"
	"sync"
)

type LoginVerificationRequest struct {
	IdToken string `json:"IdToken"`
}

func HandleLoginVerification(
	conn net.Conn,
	payload []byte,
	users map[uint32]*user.User,
	nextUserIndex *uint32,
	outgoing chan []byte,
	mu *sync.Mutex,
) bool {
	defer mu.Unlock()

	var req LoginVerificationRequest
	if err := json.Unmarshal(payload, &req); err != nil {
		println("Error parsing payload: ", err.Error())
		return false
	}

	mu.Lock()
	for _, u := range users {
		if u.Conn == conn {
			return false
		}
	}

	newUser := &user.User{
		Index: *nextUserIndex,
		ID:    43,
		Email: "sex@gmail.com",
		Conn:  conn,
	}
	*nextUserIndex++

	if _, ok := users[newUser.Index]; !ok {
		users[newUser.Index] = newUser
	}

	// Send a response packet to the client
	//resp := Message{Type: loginVerificationResponse, Payload: []byte(`{"success": true}`)}
	//outgoing <- resp.Encode()
	return true
}

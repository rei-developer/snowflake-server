package handlers

import (
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/snowflake-server/src/response"
	"github.com/snowflake-server/src/user"
)

func HandleGetUserByID(payload []byte, users map[uint32]*user.User, outgoing chan []byte, mu *sync.Mutex) {
	defer mu.Unlock()

	// parse the user ID from the payload
	userID := binary.BigEndian.Uint32(payload)

	// lock the map before accessing it
	mu.Lock()

	u, ok := users[userID]
	if !ok {
		fmt.Printf(fmt.Sprintf("User %d not found", userID))
		response.SendErrorResponse(outgoing, fmt.Sprintf("User %d not found", userID))
		return
	}

	// send the user data
	response.SendSuccessResponse(outgoing, fmt.Sprintf("User %d found with email %s", u.ID, u.Email))
}

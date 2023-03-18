package handlers

import (
	"encoding/binary"
	"fmt"
	"github.com/snowflake-server/src/response"
	"github.com/snowflake-server/src/user"
)

func HandleGetUserByID(payload []byte, outgoing chan []byte, users map[uint32]*user.User) {
	userID := binary.BigEndian.Uint32(payload)

	u, ok := users[userID]
	if !ok {
		fmt.Printf(fmt.Sprintf("User %d not found", userID))
		return
	}

	response.SendMessage(outgoing, 1, map[string]interface{}{"test": u.ID})
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/snowflake-server/src/common"
	"github.com/snowflake-server/src/user"
)

type LoginVerificationRequest struct {
	Token string `json:"token"`
}

func HandleLoginVerification(
	conn net.Conn,
	payload []byte,
	users map[uint32]*user.User,
	nextUserIndex *uint32,
	mu *sync.Mutex,
) bool {
	defer mu.Unlock()
	mu.Lock()

	var req LoginVerificationRequest
	if err := json.Unmarshal(payload, &req); err != nil {
		return false
	}

	for _, u := range users {
		if u.Conn == conn {
			return false
		}
	}

	claims, err := common.VerifyToken(req.Token)
	if err != nil || claims == nil {
		return false
	}

	userId, err := strconv.ParseUint(claims.Id, 10, 32)
	if err != nil {
		return false
	}

	newUser := &user.User{
		Model: common.Model{
			Index: *nextUserIndex,
			ID:    uint(userId),
		},
		Type:  "apple",
		UID:   "dddddd",
		Email: "지랄하네",
		Conn:  conn,
	}
	if err := user.UpsertUser(newUser); err != nil {
		fmt.Printf("Failed to create user: %v", err)
		return false
	}

	*nextUserIndex++
	users[newUser.Index] = newUser

	return true
}

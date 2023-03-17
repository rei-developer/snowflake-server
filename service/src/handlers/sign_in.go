package handlers

import (
	"encoding/json"
	"github.com/snowflake-server/src/common"
	"github.com/snowflake-server/src/user"
	"net"
)

type requestLoginVerification struct {
	Token string `json:"token"`
}

func HandleLoginVerification(
	conn net.Conn,
	payload []byte,
	users map[uint32]*user.User,
	nextUserIndex *uint32,
) bool {
	var req requestLoginVerification
	if err := json.Unmarshal(payload, &req); err != nil {
		return false
	}

	for _, u := range users {
		if u.Conn == conn {
			return false
		}
	}

	println(req.Token)

	claims, err := common.VerifyToken(req.Token)
	if err != nil || claims == nil {
		println(err.Error())
		return false
	}

	println(claims.Id)

	userModel, err := user.GetUserByUID(claims.Id)
	if err != nil {
		return false
	}

	newUser := &user.User{
		Model: common.Model{
			Index: *nextUserIndex,
			ID:    userModel.ID,
		},
		UID:    userModel.UID,
		Name:   userModel.Name,
		Sex:    userModel.Sex,
		Nation: userModel.Nation,
		Conn:   conn,
	}

	*nextUserIndex++
	users[newUser.Index] = newUser

	println(newUser.Index)
	println(newUser.ID)
	println(newUser.UID)
	println(newUser.Name)
	println(newUser.Sex)
	println(newUser.Nation)

	return true
}

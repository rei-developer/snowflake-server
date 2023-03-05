package user

import (
	"net"

	"github.com/snowflake-server/src/common"
)

type User struct {
	common.Model
	Type  string   `json:"type"`
	UID   string   `json:"uid"`
	Email string   `json:"email"`
	Conn  net.Conn `gorm:"-"`
}

const (
	userPrefix = "user:"
)

func GetUserByID(id uint) (*User, error) {
	var user User
	if err := user.GetByID(id, userPrefix, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpsertUser(user *User) error {
	if err := user.UpsertObject(user, userPrefix); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	var user User
	if err := user.DeleteObject(id, userPrefix); err != nil {
		return err
	}
	return nil
}

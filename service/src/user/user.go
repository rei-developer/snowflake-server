package user

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/snowflake-server/src/db"
	"gorm.io/gorm"
)

type User struct {
	Index   uint32     `gorm:"-"`
	ID      uint       `json:"id" gorm:"primaryKey"`
	Type    string     `json:"type"`
	UID     string     `json:"uid"`
	Email   string     `json:"email"`
	Created time.Time  `json:"created"`
	Updated time.Time  `json:"updated"`
	Deleted *time.Time `json:"deleted"`
	Conn    net.Conn   `gorm:"-"`
}

func CreateUser(user *User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUserByID(id uint) (*User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func UpdateUser(user *User) error {
	err := db.DB.Save(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func DeleteUser(id uint) error {
	err := db.DB.Delete(&User{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

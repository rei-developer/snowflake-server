package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/snowflake-server/src/db"
	"github.com/snowflake-server/src/redis"
	"gorm.io/gorm"
)

type User struct {
	Index   uint32    `gorm:"-"`
	ID      uint      `json:"id" gorm:"primaryKey"`
	Type    string    `json:"type"`
	UID     string    `json:"uid"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted time.Time `json:"deleted"`
	Conn    net.Conn  `gorm:"-"`
}

const (
	userPrefix = "user:"
)

func GetUserByID(id uint) (*User, error) {
	key := fmt.Sprintf("%s%d", userPrefix, id)

	// Try to get the user from Redis cache
	jsonStr, err := redis.Get(key)
	if err == nil {
		var user User
		if err := json.Unmarshal([]byte(jsonStr), &user); err == nil {
			return &user, nil
		}
	}

	// If the user is not in the Redis cache, get it from the database
	var user User
	err = db.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Cache the user in Redis
	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}
	if err := redis.Set(key, string(jsonBytes), time.Hour); err != nil {
		fmt.Printf("Failed to cache user %d in Redis: %v", id, err)
	}

	return &user, nil
}

func UpsertUser(user *User) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var existingUser User
		result := tx.Where("id = ?", user.ID).First(&existingUser)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to get user: %w", result.Error)
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// If the user doesn't exist, create a new user
			err := tx.Create(user).Error
			if err != nil {
				return fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			// If the user already exists, update their fields with the new values
			if err := updateUserFields(&existingUser, user); err != nil {
				return fmt.Errorf("failed to update user fields: %w", err)
			}

			// Save the updated user to the database
			if err := tx.UpdateColumns(&existingUser).Error; err != nil {
				return fmt.Errorf("failed to update user: %w", err)
			}
		}

		// Update the user in the Redis cache
		key := fmt.Sprintf("%s%d", userPrefix, user.ID)
		jsonBytes, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("failed to marshal user: %w", err)
		}
		if err := redis.Set(key, string(jsonBytes), time.Hour); err != nil {
			fmt.Printf("Failed to cache user %d in Redis: %v", user.ID, err)
		}

		return nil
	})
}

func DeleteUser(id uint) error {
	// Delete the user from the database
	err := db.DB.Delete(&User{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	// Delete the user from the Redis cache
	key := fmt.Sprintf("%s%d", userPrefix, id)
	if err := redis.Delete(key); err != nil {
		fmt.Printf("Failed to delete user %d from Redis: %v", id, err)
	}

	return nil
}

func updateUserFields(existingUser *User, newUser *User) error {
	existingValue := reflect.ValueOf(existingUser).Elem()
	newValue := reflect.ValueOf(newUser).Elem()

	for i := 0; i < existingValue.NumField(); i++ {
		existingField := existingValue.Field(i)
		newField := newValue.Field(i)
		if newField.Interface() != reflect.Zero(newField.Type()).Interface() {
			existingField.Set(newField)
		}
	}

	return nil
}

package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/snowflake-server/src/db"
	"github.com/snowflake-server/src/redis"
	"gorm.io/gorm"
)

type Model struct {
	Index uint32 `gorm:"-"`
	ID    uint   `json:"id" gorm:"primaryKey"`
}

func (m *Model) GetByID(id uint, prefix string, v interface{}) error {
	key := fmt.Sprintf("%s%d", prefix, id)

	// Try to get the object from Redis cache
	jsonStr, err := redis.Get(key)
	if err == nil {
		if err := json.Unmarshal([]byte(jsonStr), v); err == nil {
			return nil
		}
	}

	// If the object is not in the Redis cache, get it from the database
	err = db.DB.First(v, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to get object: %w", err)
	}

	// Cache the object in Redis
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %w", err)
	}
	if err := redis.Set(key, string(jsonBytes), time.Hour); err != nil {
		fmt.Printf("Failed to cache object %d in Redis: %v", id, err)
	}

	return nil
}

func (m *Model) GetByColumn(column string, value interface{}, prefix string, v interface{}) error {
	key := fmt.Sprintf("%s%s_%v", prefix, column, value)

	// Try to get the object from Redis cache
	jsonStr, err := redis.Get(key)
	if err == nil {
		if err := json.Unmarshal([]byte(jsonStr), v); err == nil {
			return nil
		}
	}

	// If the object is not in the Redis cache, get it from the database
	query := fmt.Sprintf("%s = ?", column)
	err = db.DB.Where(query, value).First(v).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to get object: %w", err)
	}

	// Cache the object in Redis
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %w", err)
	}
	if err := redis.Set(key, string(jsonBytes), time.Hour); err != nil {
		fmt.Printf("Failed to cache object %v in Redis: %v", value, err)
	}

	return nil
}

func (m *Model) UpdateFields(existingObj interface{}, newObj interface{}) {
	vExistingObj := reflect.ValueOf(existingObj).Elem()
	vNewObj := reflect.ValueOf(newObj).Elem()

	for i := 0; i < vExistingObj.NumField(); i++ {
		existingFieldValue := vExistingObj.Field(i)
		newFieldValue := vNewObj.Field(i)
		if !newFieldValue.IsZero() {
			existingFieldValue.Set(newFieldValue)
		}
	}
}

func (m *Model) UpsertObject(obj interface{}, prefix string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		id := reflect.ValueOf(obj).Elem().FieldByName("ID").Uint()

		// Get the existing object from the database
		existingObj := reflect.New(reflect.TypeOf(obj).Elem()).Interface()
		result := tx.Where("id = ?", id).First(existingObj)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to get object: %w", result.Error)
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// If the object doesn't exist, create a new object
			err := tx.Create(obj).Error
			if err != nil {
				return fmt.Errorf("failed to create object: %w", err)
			}
		} else {
			// If the object already exists, update its fields with the new values
			m.UpdateFields(existingObj, obj)

			// Save the updated object to the database
			if err := tx.Model(obj).UpdateColumns(existingObj).Error; err != nil {
				return fmt.Errorf("failed to update object: %w", err)
			}
		}

		// Update the object in the Redis cache
		key := fmt.Sprintf("%s%d", prefix, id)
		jsonBytes, err := json.Marshal(obj)
		if err != nil {
			return fmt.Errorf("failed to marshal object: %w", err)
		}
		if err := redis.Set(key, string(jsonBytes), time.Hour); err != nil {
			fmt.Printf("Failed to cache object %d in Redis: %v", id, err)
		}

		return nil
	})
}

func (m *Model) DeleteObject(id uint, prefix string) error {
	tx := db.DB.Begin()

	// Delete the object from the database
	if err := tx.Delete(reflect.New(reflect.TypeOf(m).Elem()).Interface(), id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%s not found", prefix)
		}
		return fmt.Errorf("failed to delete %s: %w", prefix, err)
	}

	// Delete the object from the Redis cache
	key := fmt.Sprintf("%s%d", prefix, id)
	if err := redis.Delete(key); err != nil {
		fmt.Printf("Failed to delete %s %d from Redis: %v", prefix, id, err)
	}

	return tx.Commit().Error
}

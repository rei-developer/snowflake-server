package db

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

type Config struct {
	DB struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		DBName       string `yaml:"dbname"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
	} `yaml:"db"`
}

var DB *gorm.DB

func Connect() error {
	// Read the configuration file
	configData, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %w", err)
	}

	// Parse the configuration
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return fmt.Errorf("failed to parse configuration: %w", err)
	}

	// Set up the database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set the maximum number of open connections
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB instance: %w", err)
	}

	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)

	// Set the maximum number of idle connections
	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)

	// Test the connection
	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

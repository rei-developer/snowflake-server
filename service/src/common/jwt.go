package common

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
)

type jwtConfig struct {
	JWT struct {
		DefaultStrategy string `yaml:"defaultStrategy"`
		SecretKey       string `yaml:"secretKey"`
	} `yaml:"jwt"`
}

func VerifyToken(tokenString string) (*jwt.StandardClaims, error) {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	secretKey := []byte(config.JWT.SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func loadConfig() (*jwtConfig, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config.yaml")

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config jwtConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

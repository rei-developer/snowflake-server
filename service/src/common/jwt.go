package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	JWT struct {
		DefaultStrategy string `yaml:"defaultStrategy"`
		SecretKey       string `yaml:"secretKey"`
	} `yaml:"jwt"`
}

func VerifyToken(tokenString string) (*jwt.StandardClaims, error) {
	// Read the configuration file
	configData, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	// Parse the configuration
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
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

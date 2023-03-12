package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type openAiConfig struct {
	OpenAI struct {
		SecretKey string `yaml:"secretKey"`
	} `yaml:"openai"`
}

type chatRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type chatResponse struct {
	Choices []struct {
		Index        int     `json:"index"`
		FinishReason *string `json:"finish_reason,omitempty"`
		Message      struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func SendChatRequest(content string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	url := "https://api.openai.com/v1/chat/completions"

	requestBody := chatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		println(err.Error())
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		println(err.Error())
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.OpenAI.SecretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return "", errors.New("Chat request failed with status code " + strconv.Itoa(resp.StatusCode))
	}

	var response chatResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		println(err.Error())
		return "", err
	}

	return strings.TrimSpace(response.Choices[0].Message.Content), nil
}

func loadConfig() (*openAiConfig, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config.yaml")

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config openAiConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

package novelai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/snowflake-server/src/common"
	"github.com/snowflake-server/src/generated_image"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

type novelAiConfig struct {
	S3 struct {
		KeyNamePrefix string `yaml:"keyNamePrefix"`
	} `yaml:"s3"`
	NovelAI struct {
		ApiKey string `yaml:"apiKey"`
	} `yaml:"novelai"`
}

type LoginRequest struct {
	Key string `json:"key"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type ImageRequest struct {
	Input      string                 `json:"input"`
	Model      string                 `json:"model"`
	Parameters map[string]interface{} `json:"parameters"`
}

var accessToken string = ""

func GenerateImage(input string) string {
	prompt, hash := common.GetPromptHash(input)

	image, err := generated_image.GetGeneratedImageByHash(hash)
	if err != nil {
		println(err.Error())
		return ""
	}
	if image.ID > 0 {
		println("이미 있군요...")
		return image.Hash
	}

	request := ImageRequest{
		Input: prompt,
		Model: "safe-diffusion",
		Parameters: map[string]interface{}{
			"width":         512,
			"height":        768,
			"scale":         11,
			"seed":          3698792792,
			"sampler":       "k_euler_ancestral",
			"steps":         28,
			"n_samples":     1,
			"ucPreset":      0,
			"qualityToggle": true,
			"autoSmea":      true,
			"uc":            "nsfw, lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, face, human",
		},
	}

	config, authToken, err := getAuthToken()

	response, err := generateImage(request, authToken)
	if err != nil {
		panic(err)
	}

	dataRegex := regexp.MustCompile(`data:(.*)`)
	dataMatch := dataRegex.FindStringSubmatch(response)
	if len(dataMatch) > 1 {
		data := dataMatch[1]

		err := common.UploadToS3(data, config.S3.KeyNamePrefix+hash)
		if err != nil {
			println(err.Error())
			return ""
		}

		newImage := &generated_image.GeneratedImage{
			UserID: 0,
			Prompt: prompt,
			Hash:   hash,
		}
		if err := generated_image.UpsertGeneratedImage(newImage); err != nil {
			fmt.Printf("Failed to create user: %v", err)
		}

		return hash
	} else {
		fmt.Println("No match found")
	}

	return ""
}

func loadConfig() (*novelAiConfig, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config.yaml")

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config novelAiConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getAuthToken() (*novelAiConfig, string, error) {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	if accessToken != "" {
		return config, accessToken, nil
	}

	// Create a new HTTP request
	reqBody := &LoginRequest{Key: config.NovelAI.ApiKey}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return config, "", err
	}
	req, err := http.NewRequest("POST", "https://api.novelai.net/user/login", bytes.NewBuffer(reqBytes))
	if err != nil {
		return config, "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send the HTTP request and retrieve the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return config, "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)

	// Read the response body and parse the JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return config, "", err
	}
	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return config, "", err
	}

	// Store the access token and return it
	accessToken = loginResp.AccessToken
	return config, accessToken, nil
}

func generateImage(request ImageRequest, authToken string) (string, error) {
	url := "https://api.novelai.net/ai/generate-image"

	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func saveImage(base64Data string, filename string, folder string) error {
	// Decode the base64 data
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Create the file path
	filepath := path.Join(wd, folder, filename)

	// Create the folder if it doesn't exist
	err = os.MkdirAll(path.Dir(filepath), os.ModePerm)
	if err != nil {
		return err
	}

	// Create a new file to write the image data
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println(err)
		}
	}(file)

	// Write the image data to the file
	n, err := file.Write(imageData)
	if err != nil {
		return err
	}
	fmt.Println("Bytes written:", n)

	return nil
}

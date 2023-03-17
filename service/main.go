package main

import (
	"bufio"
	"fmt"
	"github.com/snowflake-server/src/db"
	"github.com/snowflake-server/src/openai"
	"github.com/snowflake-server/src/redis"
	"github.com/snowflake-server/src/server"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type serverConfig struct {
	Service struct {
		Port string `yaml:"port"`
	} `yaml:"service"`
}

var validCommands map[string]string

func init() {
	validCommands = map[string]string{
		"talk": "Talk to ChatGPT AI.",
		"help": "Display this help message.",
		"exit": "Exit the program.",
	}
}

func main() {
	regexp.MustCompile("")

	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	s, err := server.NewServer(config.Service.Port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Printf("Server started on port %s. Type 'help' for a list of commands. Press enter to exit.\n", config.Service.Port)

	err = db.Connect()
	if err != nil {
		panic(err)
	}

	err = redis.Connect()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		if err := handleCommand(text, validCommands, reader); err != nil {
			fmt.Println(err)
		}
		if text == "exit" {
			if confirmExit(reader) {
				break
			}
		}
	}
}

func loadConfig() (*serverConfig, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config.yaml")

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config serverConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func handleCommand(cmd string, validCommands map[string]string, reader *bufio.Reader) error {
	parts := strings.SplitN(cmd, " ", 2)
	command := parts[0]
	content := ""
	if len(parts) > 1 {
		content = parts[1]
	}

	switch command {
	case "talk":
		resultChan := make(chan string)
		errChan := make(chan error)

		go func() {
			println("Me : " + content)
			result, err := openai.SendChatRequest(content)
			if err != nil {
				errChan <- err
			} else {
				resultChan <- result
			}
		}()

		select {
		case result := <-resultChan:
			fmt.Println("AI: " + result)
		case err := <-errChan:
			fmt.Println(err.Error())
		}

		return nil
	case "help":
		fmt.Println("Valid commands:")
		for cmd, usage := range validCommands {
			fmt.Printf("%-10s%s\n", cmd, usage)
		}
	case "exit":
		return nil
	default:
		if usage, ok := validCommands[cmd]; ok {
			fmt.Printf("%s (%s)\n", cmd, usage)
		} else {
			for name, _ := range validCommands {
				if strings.HasPrefix(name, cmd) {
					fmt.Printf("Did you mean '%s'? (%s)\n", name, validCommands[name])
					return nil
				}
			}
			return fmt.Errorf("unknown command: %s", cmd)
		}
	}
	return nil
}

func confirmExit(reader *bufio.Reader) bool {
	fmt.Print("Are you sure you want to exit? (yes/no) ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	return answer == "yes" || answer == "y" || answer == "ok"
}

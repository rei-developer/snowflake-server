package main

import (
	"bufio"
	"fmt"
	"github.com/snowflake-server/src/db"
	"github.com/snowflake-server/src/redis"
	"github.com/snowflake-server/src/server"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Config struct {
	Service struct {
		Port string `yaml:"port"`
	} `yaml:"service"`
}

var validCommands map[string]string

func init() {
	validCommands = map[string]string{
		"help": "Display this help message.",
		"exit": "Exit the program.",
	}
}

func main() {
	configFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
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

func handleCommand(cmd string, validCommands map[string]string, reader *bufio.Reader) error {
	switch cmd {
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

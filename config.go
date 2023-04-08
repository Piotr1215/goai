package goai

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Model         string  `yaml:"model"`
	Temperature   float64 `yaml:"temperature"`
	MaxTokens     int     `yaml:"max_tokens"`
	Safety        bool    `yaml:"safety"`
	OpenAI_APIKey string  `yaml:"openai_api_key"`
}

func configFilePath(filename string) string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		homeDir = usr.HomeDir
	}

	configDir := filepath.Join(homeDir, ".config", "goai")
	return filepath.Join(configDir, filename)
}

func parseConfig(configContent string) Config {
	var config Config
	err := yaml.Unmarshal([]byte(configContent), &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

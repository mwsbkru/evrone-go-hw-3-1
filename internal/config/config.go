package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port     string
	Host     string
	Greeting string
}

const PORT_ENV_NAME = "port"
const HOST_ENV_NAME = "host"
const GREETING_ENV_NAME = "greeting"

func NewConfig() (Config, error) {
	missingParams := make([]string, 0, 3)

	port, ok := os.LookupEnv(PORT_ENV_NAME)
	if !ok {
		missingParams = append(missingParams, PORT_ENV_NAME)
	}

	host, ok := os.LookupEnv(HOST_ENV_NAME)
	if !ok {
		missingParams = append(missingParams, HOST_ENV_NAME)
	}

	greeting, ok := os.LookupEnv(GREETING_ENV_NAME)
	if !ok {
		missingParams = append(missingParams, GREETING_ENV_NAME)
	}

	if len(missingParams) != 0 {
		missingFieldsStr := strings.Join(missingParams, ", ")
		err := fmt.Errorf("не установлены значения env-переменных: %s", missingFieldsStr)
		return Config{}, err
	}

	return Config{Port: port, Host: host, Greeting: greeting}, nil
}

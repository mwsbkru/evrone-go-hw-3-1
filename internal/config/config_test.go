package config

import (
	"errors"
	"os"
	"testing"
)

type testNewConfigCase struct {
	name           string
	expectedConfig Config
	expectedError  error
	initEnvFunc    func()
}

const CORRECT_PORT = "80"
const CORRECT_HOST = "localhost"
const CORRECT_GREETING = "TestGreet"

func TestNewConfig(t *testing.T) {
	testCases := []testNewConfigCase{
		{
			name:          "without any env var",
			expectedError: errors.New("не установлены значения env-переменных: port, host, greeting"),
			initEnvFunc: func() {
				os.Unsetenv(PORT_ENV_NAME)
				os.Unsetenv(HOST_ENV_NAME)
				os.Unsetenv(GREETING_ENV_NAME)
			},
		},
		{
			name: "without correctly initialized env var",
			expectedConfig: Config{
				Port:     CORRECT_PORT,
				Host:     CORRECT_HOST,
				Greeting: CORRECT_GREETING,
			},
			initEnvFunc: func() {
				os.Setenv(PORT_ENV_NAME, CORRECT_PORT)
				os.Setenv(HOST_ENV_NAME, CORRECT_HOST)
				os.Setenv(GREETING_ENV_NAME, CORRECT_GREETING)
			},
		},
	}

	for _, tc := range testCases {
		tc.initEnvFunc()

		cfg, err := NewConfig()
		if tc.expectedError != nil && err == nil {
			t.Error("Expected to return error, but no error returned")
			continue
		}

		if cfg != tc.expectedConfig {
			t.Errorf("Wrong config: expected %v; actual: %v", tc.expectedConfig, cfg)
		}
	}
}

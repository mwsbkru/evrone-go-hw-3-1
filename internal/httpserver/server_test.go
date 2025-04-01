package httpserver

import (
	"encoding/json"
	"hw_3_1/internal/config"
	"hw_3_1/pkg/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHello(t *testing.T) {
	request, err := http.NewRequest("GET", "/greet", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()

	cfg := config.Config{
		Greeting: "TestGreet",
	}

	params := httpserver.GetHelloParams{Name: "Tester"}

	server := NewServer(&cfg)
	server.GetHello(responseRecorder, request, params)

	var result httpserver.GetHelloResponse

	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	expectedGreet := "TestGreet Tester!"
	if result.Greeting != expectedGreet {
		t.Errorf("Wrong response body: expected: %s; actual: %s", expectedGreet, result.Greeting)
	}
}

package httpserver

import (
	"encoding/json"
	"fmt"
	"hw_3_1/internal/config"
	"hw_3_1/pkg/httpserver"
	"log/slog"
	"net/http"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) httpserver.ServerInterface {
	return Server{cfg: cfg}
}

func (s Server) GetHello(w http.ResponseWriter, r *http.Request, params httpserver.GetHelloParams) {
	greetingStr := fmt.Sprintf("%s %s!", s.cfg.Greeting, params.Name)
	response := httpserver.GetHelloResponse{Greeting: greetingStr}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Не удалось сериализовать тело ответа в json", slog.String("error", err.Error()))
		return
	}

	w.Write(jsonResponse)
}

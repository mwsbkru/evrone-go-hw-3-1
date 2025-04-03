package main

import (
	"fmt"
	"hw_3_1/internal/config"
	"hw_3_1/internal/httpserver"
	httpserver2 "hw_3_1/pkg/httpserver"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Не удалось прочитать конфиг приложения", slog.String("error", err.Error()))
		os.Exit(1)
	}

	server := httpserver.NewServer(&cfg)
	router := http.NewServeMux()
	handler := httpserver2.HandlerFromMux(server, router)
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	httpServer := &http.Server{
		Handler: handler,
		Addr:    addr,
	}

	slog.Info("Запущен сервер", slog.String("address", addr))
	err = httpServer.ListenAndServe()
	if err != nil {
		slog.Error("Не удалось запустить сервер", slog.String("error", err.Error()))
	}
}

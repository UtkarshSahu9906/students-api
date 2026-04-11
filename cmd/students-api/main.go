package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UtkarshSahu9906/students-api/internal/config"
	"github.com/UtkarshSahu9906/students-api/internal/http/handlers/student"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

router.HandleFunc("POST /api/students", student.New())
	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Printf("Server started %s\n", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}
}

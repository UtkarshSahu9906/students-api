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
	"github.com/UtkarshSahu9906/students-api/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %s", err.Error())
	}

	slog.Info("storate initializad", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	router := http.NewServeMux()




	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetbyID(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	// router.HandleFunc("PUT /api/students/{id}", student.Update(storage))
	// router.HandleFunc("DELETE /api/students/{id}", student.Delete(storage))
	// router.HandleFunc("GET /api/students/", student.GetAll(storage))








	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Printf("Server started %s\n", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	err = server.ListenAndServe()
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

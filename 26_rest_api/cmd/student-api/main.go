package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itz-prashant/student-api/internal/config"
	"github.com/itz-prashant/student-api/internal/http/handlers/student"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// databse setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	srver := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started",slog.String("Address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srver.ListenAndServe()

		if err != nil  && err != http.ErrServerClosed {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srver.Shutdown(ctx)
	if err != nil  {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}

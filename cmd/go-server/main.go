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

	"github.com/Sudip34559/go-server/internal/config"
)

func main() {
	conf := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server := http.Server{
		Addr:    conf.HTTPServer.Address,
		Handler: router,
	}

	slog.Info("server running", slog.String("address:", conf.HTTPServer.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("faied to shutdown server", slog.String("error: ", err.Error()))
	}

	slog.Info("server shutdown successfully")
}

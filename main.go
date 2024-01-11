package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/programkingstar/task-management-go.git/api"
)

func initServer(ctx context.Context) error {
	server := api.NewServer()
	return server.Start(ctx)
}

// @version 1
// @title Notethingness API
// @description This is a sample server for Notethingness API.

// @host localhost:3000
// @BasePath /api
func main() {
	serverCtx, stopCtx := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopCtx()

	err := initServer(serverCtx)
	if err != nil {
		log.Fatal("can't run the server: ", err)
		slog.Error(" [ 💢Cannot run the server! ] " + "\nError: " + err.Error())
		slog.ErrorContext(serverCtx, " [ 💢Cannot run the server! ] "+"\nError: "+err.Error())
		os.Exit(1)
	}
}

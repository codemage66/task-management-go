package api

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/programkingstar/task-management-go.git/db"
	"github.com/programkingstar/task-management-go.git/types"
	"github.com/programkingstar/task-management-go.git/util"
)

type Server struct {
	router http.Handler
	db     *sql.DB
	config types.Env
}

func NewServer() *Server {
	config := util.GetEnv()

	db := db.NewDatabase()
	slog.Info("[ ☘️ Connect to DB POSTGRES ]")

	store, err := db.Connect(config.DBUrl)
	if err != nil {
		panic(err)
	}

	server := &Server{
		db:     store,
		config: config,
	}

	slog.Info("[ ☘️ Run migration rollback ]")
	if err := db.RollbackMigration(config.DBUrl, "file://db/migration"); err != nil {
		panic(err)
	}

	slog.Info("[ ☘️ Run migration ]")
	if err := db.RunMigration(config.DBUrl, "file://db/migration"); err != nil {
		panic(err)
	}

	server.Router()
	return server
}

// Start starts the server
func (s *Server) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.config.Port),
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("[ See Documentation on http://localhost:" + s.config.Port + "/openapi ]")
	slog.Info("[ Server started on port: " + s.config.Port + " ]")
	defer func() {}()

	// Using a buffered channel to avoid goroutine leaks
	channel := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			channel <- error(err)
		}
		close(channel)
	}()

	select {
	case err := <-channel:
		slog.Error("[ Failed to start server ]" + "\nError: " + err.Error())
		return err
	case <-ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			return error(err)
		}
		return nil
	}
}

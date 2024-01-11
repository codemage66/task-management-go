package main

import (
	"context"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestInitServer(t *testing.T) {
	if err := os.Setenv("PORT", "3000"); err != nil {
		t.Fatal("Failed to set environment variable for testing:", err)
	}
	defer os.Unsetenv("PORT")

	testCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)

	go func() {
		main()
	}()

	select {
	case <-testCtx.Done():
		return
	case <-interruptCh:
		return
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

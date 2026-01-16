package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flgs := getFlags()
	// First signal → graceful shutdown
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Second signal → force exit
	forceExit := make(chan os.Signal, 1)
	signal.Notify(forceExit, os.Interrupt, syscall.SIGTERM)

	c, cl := NewContainer(flgs)

	log.Println("serverpulse agent started")

	// Start agent
	go func(ctx context.Context) {
		err := c.agent.start(ctx)
		if err != nil {
			log.Printf("agent start returned error: %v", err)
		}
	}(ctx)

	// Wait for first signal
	<-ctx.Done()
	log.Println("shutdown signal received (graceful)")

	// Graceful shutdown window
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	// Listen for second signal
	go func() {
		<-forceExit
		log.Println("second signal received, forcing exit")
		os.Exit(1)
	}()

	// Run cleanup
	cl(shutdownCtx)

	log.Println("agent exited cleanly")
}

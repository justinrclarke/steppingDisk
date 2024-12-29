package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"steppingDisk/internal/config"
	"steppingDisk/internal/server"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	srv, err := server.New(cfg)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	go func() {
		if err := srv.Start(ctx); err != nil {
			log.Printf("Server error: %v", err)
			cancel()
		}
	}()

	<-sigChan
	log.Println("Shutting down gracefully...")
	cancel()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}

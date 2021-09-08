package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"simplesurance.com/pkg/config"
	"simplesurance.com/pkg/server"
	"syscall"
)

func main() {
	logger := log.New(os.Stdout, "app : ", log.LstdFlags)

	if err := run(logger); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	conf := config.Config()

	logger.Println("Initializing application")
	defer logger.Println("Application stopped")

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	srv := server.New(conf, logger)

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- srv.Serve()
	}()

	select {
	case err := <-serverErrors:
		return err

	case <-shutdownCtx.Done():
		logger.Println("Starting graceful shutdown")

		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
		defer timeoutCancel()

		if err := srv.Stop(timeoutCtx); err != nil {
			return fmt.Errorf("failed to shutdown app: %w", err)
		}
	}

	return nil
}

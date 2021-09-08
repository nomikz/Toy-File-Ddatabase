package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"simplesurance.com/pkg/config"
	"simplesurance.com/pkg/db"
	"simplesurance.com/pkg/handlers"
)

func main() {
	logger := log.New(os.Stdout, "app : ", log.LstdFlags)

	if err := run(logger); err != nil {
		logger.Println("error: ", err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	conf := config.Config()

	logger.Println("Initializing the application")
	defer logger.Println("Application stopped")

	ds, err := db.New(conf.DataStoreFileName)
	if err != nil {
		return err
	}

	// catch shutdown signal for graceful shutdown
	shutdownCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	api := http.Server{
		Addr:         conf.ListenAddr,
		Handler:      handlers.Routes(logger, ds, conf.SecondsCount),
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Println("Serving on port", conf.ListenAddr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err = <-serverErrors:
		return err

	case <-shutdownCtx.Done():
		logger.Println("Starting graceful shutdown")

		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
		defer timeoutCancel()

		if err = api.Shutdown(timeoutCtx); err != nil {
			log.Println("Failed to shutdown gracefully")

			if err := api.Close(); err != nil {
				return fmt.Errorf("failed to close server: %w", err)
			}
		}
	}

	return nil
}

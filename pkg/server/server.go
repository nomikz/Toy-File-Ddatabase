package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"simplesurance.com/pkg/config"
	"simplesurance.com/pkg/server/handlers"
)

type Server struct {
	logger       *log.Logger
	listenAddr   string
	readTimeout  time.Duration
	writeTimeout time.Duration
	server       *http.Server
}

func New(conf config.Conf, logger *log.Logger) *Server {
	return &Server{
		logger:       logger,
		listenAddr:   conf.ListenAddr,
		readTimeout:  conf.ReadTimeout,
		writeTimeout: conf.WriteTimeout,
	}
}

func (srv *Server) Serve() error {
	srv.server = &http.Server{
		Addr:         srv.listenAddr,
		Handler:      handlers.Routes(srv.logger),
		ReadTimeout:  srv.readTimeout,
		WriteTimeout: srv.writeTimeout,
	}

	srv.logger.Println("Serving on port", srv.listenAddr)
	return srv.server.ListenAndServe()
}

func (srv *Server) Stop(ctx context.Context) error {
	if err := srv.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		srv.logger.Println("Failed to gracefully shutdown server", err)
		return srv.server.Close()
	}

	return nil
}

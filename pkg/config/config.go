package config

import (
	"flag"
	"time"
)

type Conf struct {
	ListenAddr      string
	WriteTimeout    time.Duration
	ReadTimeout     time.Duration
	ShutdownTimeout time.Duration
}

// Config parses env variable and return config struct
func Config() Conf {
	var listenAddr string
	var writeTimeout int
	var readTimeout int
	var shutdownTimeout int

	flag.StringVar(&listenAddr, "listen-addr", ":8081", "server listen address")
	flag.IntVar(&writeTimeout, "http-read-timeout", 5, "http read timeout (second)")
	flag.IntVar(&readTimeout, "http-write-timeout", 5, "http write timeout (second)")
	flag.IntVar(&shutdownTimeout, "server-shutdown-timeout", 5, "http shutdown timeout (second)")

	flag.Parse()

	return Conf{
		ListenAddr:      listenAddr,
		WriteTimeout:    time.Second * time.Duration(writeTimeout),
		ReadTimeout:     time.Second * time.Duration(readTimeout),
		ShutdownTimeout: time.Second * time.Duration(shutdownTimeout),
	}
}

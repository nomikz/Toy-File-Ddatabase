package config

import (
	"flag"
	"time"
)

type Conf struct {
	ListenAddr        string
	DataStoreFileName string
	SecondsCount      int
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	ShutdownTimeout   time.Duration
}

// Config parses env variable and return config struct
func Config() Conf {
	var listenAddr string
	var dataStoreFileName string
	var secondsCount int
	var writeTimeout int
	var readTimeout int
	var shutdownTimeout int

	flag.StringVar(&listenAddr, "listen-addr", ":8081", "server listen address")
	flag.StringVar(&dataStoreFileName, "ds-file-name", "db.json", "filename. if not it get created")
	flag.IntVar(&secondsCount, "seconds-count", 60, "store request count of the last N seconds")
	flag.IntVar(&writeTimeout, "http-read-timeout", 5, "http read timeout (second)")
	flag.IntVar(&readTimeout, "http-write-timeout", 5, "http write timeout (second)")
	flag.IntVar(&shutdownTimeout, "server-shutdown-timeout", 5, "http shutdown timeout (second)")

	flag.Parse()

	return Conf{
		ListenAddr:        listenAddr,
		DataStoreFileName: dataStoreFileName,
		SecondsCount:      secondsCount,
		WriteTimeout:      time.Second * time.Duration(writeTimeout),
		ReadTimeout:       time.Second * time.Duration(readTimeout),
		ShutdownTimeout:   time.Second * time.Duration(shutdownTimeout),
	}
}

package handlers

import (
	"log"
	"net/http"

	"simplesurance.com/pkg/db"
)

func Routes(log *log.Logger, db *db.DataStore, secCount int) http.Handler {
	mux := http.NewServeMux()

	c := counter{
		secCount: secCount,
		db:       db,
		logger:   log,
	}

	mux.Handle("/", http.HandlerFunc(c.counter))

	return mux
}

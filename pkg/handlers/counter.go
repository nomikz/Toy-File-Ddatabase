package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"simplesurance.com/pkg/db"
)

type counter struct {
	secCount int

	db     *db.DataStore
	logger *log.Logger
}

func (c counter) counter(w http.ResponseWriter, r *http.Request) {
	count, err := c.db.Counter(c.secCount)

	if errors.Is(err, db.ErrFileNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprintf(w, "Database file not found. Please create file.")
		return
	}
	if err != nil {
		c.logger.Println("[ERROR] counter: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Counter: %v", count)
	if err != nil {
		c.logger.Println("[ERROR] counter: ", err)
	}
}

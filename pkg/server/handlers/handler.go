package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Routes(log *log.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileName := "db.json"
		if _, err := ioutil.ReadFile(fileName); err != nil {
			str := "[]"
			if err = ioutil.WriteFile(fileName, []byte(str), 0644); err != nil {
				log.Fatal(err)
			}
		}
	}))

	return mux
}
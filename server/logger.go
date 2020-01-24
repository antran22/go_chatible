package server

import (
	"log"
	"net/http"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request at", r.URL)
		handler.ServeHTTP(w, r)
	})
}

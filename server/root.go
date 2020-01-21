package server

import (
	"log"
	"net/http"
	"os"
)

func Serve() {
	addr := os.Getenv("HOSTNAME") + ":" + os.Getenv("PORT")
	log.Println("Listening on address:", addr)
	log.Fatal(http.ListenAndServe(addr, http.DefaultServeMux))
}

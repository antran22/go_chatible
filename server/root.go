package server

import (
	"log"
	"net/http"
	"os"

	"go_chatible/api"
	"go_chatible/server/webhook"
)

func pingHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("Pong")); err != nil {
		log.Println(err)
	}
}

func Serve() {
	server := http.NewServeMux()
	api.InitAPIs()
	server.HandleFunc("/ping", pingHandler)
	server.HandleFunc("/webhook", webhook.Handler)
	addr := os.Getenv("HOSTNAME") + ":" + os.Getenv("PORT")
	log.Println("Listening on address:", addr)
	log.Fatalln(http.ListenAndServe(addr, logRequest(server)))
}

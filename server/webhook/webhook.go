package webhook

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	messageController "go_chatible/controller/message"
	"go_chatible/model/message"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		webhookMessageHandler(w, r)
	} else if r.Method == "GET" {
		webhookRegisterHandler(w, r)
	}
}

func webhookRegisterHandler(w http.ResponseWriter, r *http.Request) {
	verifyToken := os.Getenv("WEBHOOK_TOKEN")
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")
	if mode == "subscribe" && token == verifyToken {
		if _, err := w.Write([]byte(challenge)); err != nil {
			log.Println("Error", err)
		}
	} else {
		w.WriteHeader(400)
		if _, err := w.Write([]byte("Nope")); err != nil {
			log.Println("Error", err)
		}
	}
}

func webhookMessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	messages := message.DecodeMessage(body)
	if err := r.ParseForm(); err != nil {
		log.Println("Error", err)
	}
	for _, msg := range messages {
		messageController.ProcessMessage(msg)
	}
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Println("Error", err)
	}
}

package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go_chatible/model/message"
)

type MessengerWorkerPool struct {
	url        string
	msgChannel chan message.Message
}

func sendMessage(client *http.Client, url string, msg message.Message) (*http.Response, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return &http.Response{}, err
	}
	bodyReader := bytes.NewReader(body)
	log.Println(string(body))
	resp, err := client.Post(url, "application/json", bodyReader)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}

func (pool MessengerWorkerPool) SpawnWorker() {
	client := http.Client{}
	for msg := range pool.msgChannel {
		resp, err := sendMessage(&client, pool.url, msg)
		if err != nil {
			log.Println("Error", err)
		}
		log.Println(resp)
	}

}

func (pool MessengerWorkerPool) QueueMessage(msg message.Message) {
	pool.msgChannel <- msg
}

func NewWorkerPool(size int) *MessengerWorkerPool {
	res := &MessengerWorkerPool{
		url:        "https://graph.facebook.com/v5.0/me/messages?access_token=" + os.Getenv("PAGE_TOKEN"),
		msgChannel: make(chan message.Message),
	}
	for w := 0; w < size; w++ {
		go res.SpawnWorker()
	}
	return res
}

package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go_chatible/model/message"
)

type MessengerSenderPool struct {
	url        string
	msgChannel chan message.Message
}

func sendMessage(client *http.Client, url string, msg message.Message) (*http.Response, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return &http.Response{}, err
	}
	bodyReader := bytes.NewReader(body)
	resp, err := client.Post(url, "application/json", bodyReader)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}

func (pool MessengerSenderPool) SpawnWorker(id int) {
	client := http.Client{}
	for msg := range pool.msgChannel {
		resp, err := sendMessage(&client, pool.url, msg)
		if err != nil {
			log.Println("Worker", id, "Messenger API Error", err)
		} else if resp.StatusCode != 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error", err)
			}
			log.Println("Worker", id, "Messenger API failed, response:", string(body))
		}
	}

}

func (pool MessengerSenderPool) QueueMessage(msg message.Message) {
	pool.msgChannel <- msg
}

func NewWorkerPool(size int) *MessengerSenderPool {
	res := &MessengerSenderPool{
		url:        "https://graph.facebook.com/v5.0/me/messages?access_token=" + os.Getenv("PAGE_ACCESS_TOKEN"),
		msgChannel: make(chan message.Message),
	}
	for w := 0; w < size; w++ {
		go res.SpawnWorker(w)
	}
	return res
}

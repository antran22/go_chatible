package message

import (
	"encoding/json"
	"log"
)

func DecodeMessage(jsonBody []byte) (res []Message) {
	data := Data{}
	res = make([]Message, 0)
	if err := json.Unmarshal(jsonBody, &data); err != nil {
		log.Println("Error", err)
	}
	for _, entry := range data.Entry {
		for _, message := range entry.Messages {
			res = append(res, *message)
		}
	}
	return
}

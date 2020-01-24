package message

import (
	"go_chatible/api"
	"go_chatible/model/message"
)

func ProcessMessage(msg message.Message) {
	if msg.Postback != nil {
		ProcessPostback(msg)
	} else {
		msgs := message.BuildForwardMessage(msg, msg.Sender.ID)
		for _, outMsg := range msgs {
			api.MessengerWorker.QueueMessage(outMsg)
		}
	}
}

package message

import (
	"log"

	"go_chatible/api"
	"go_chatible/controller/user"
	"go_chatible/model/message"
)

func ProcessMessage(msg message.Message) {
	err := user.FetchUser(msg.Sender)
	if err != nil {
		log.Println(err)
		// Send error message back to user
	}
	if msg.Postback != nil {
		ProcessPostback(msg)
	} else {
		msgs := message.BuildForwardMessage(msg, msg.Sender.ID)
		for _, outMsg := range msgs {
			_ = api.MessengerSender.SendMessage(outMsg)
		}
	}
}

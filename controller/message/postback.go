package message

import (
	"go_chatible/model/message"
)

func ProcessPostback(msg message.Message) {
	switch message.PostbackType(msg.Postback.PostbackPayload) {
	case message.Start:
		//msg.Sender.Start()
	case message.RequestEnd:
		//msg.Sender.RequestEnd()
	case message.ConfirmEnd:
		//msg.Sender.ConfirmEnd()
	case message.ChangePreferrence:
		//msg.Sender.ChangePreferrence()
	case message.PreferFemale:
		//msg.Sender.SetPreferrence("female")
	case message.PreferMale:
		//msg.Sender.SetPreferrence("male")
	}
}

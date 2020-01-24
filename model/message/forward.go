package message

import (
	"go_chatible/model/user"
)

func BuildForwardMessage(message Message, recipientId string) []Message {
	res := make([]Message, 0)
	if message.Content.Text != "" {
		res = append(res, Message{
			Content: &Content{
				Text: message.Content.Text,
			},
			Recipient: &user.User{ID: recipientId},
		})
	}
	for _, atm := range message.Content.Attachments {
		msg := Message{
			Sender:    nil,
			Recipient: &user.User{ID: recipientId},
			Content: &Content{
				Attachment: atm,
			},
		}
		res = append(res, msg)
	}
	return res
}

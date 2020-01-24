package message

import (
	"go_chatible/model/user"
)

type AttachmentPayload struct {
	URL      string             `json:"url,omitempty"`
	Type     string             `json:"template_type,omitempty"`
	Elements []*templateElement `json:"elements,omitempty"`
}

type Attachment struct {
	Type    string             `json:"type,omitempty"`
	Payload *AttachmentPayload `json:"payload,omitempty"`
}

type Content struct {
	Text        string        `json:"text,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
	Attachment  *Attachment   `json:"attachment,omitempty"`
}
type Message struct {
	Sender    *user.User `json:"sender,omitempty"`
	Recipient *user.User `json:"recipient,omitempty"`
	Timestamp int        `json:"timestamp,omitempty"`
	Content   *Content   `json:"message,omitempty"`
	Postback  *Postback  `json:"postback,omitempty"`
}

type Data struct {
	Entry []struct {
		ID       string     `json:"id,omitempty"`
		Time     int        `json:"time,omitempty"`
		Messages []*Message `json:"messaging,omitempty"`
	} `json:"entry,omitempty"`
}

package message

type templateElement struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle,omitempty"`
	Image    string    `json:"image_url,omitempty"`
	Buttons  []*Button `json:"buttons,omitempty"`
}

type Button struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
}

func MakeButton(tp PostbackType) *Button {
	pb := MakePostback(tp)
	return &Button{
		Type:    "postback",
		Title:   pb.Title,
		Payload: pb.PostbackPayload,
	}
}

func MakeTemplateMessage(title string, subtitle string, image string, buttons []*Button) Message {
	elm := templateElement{
		Title:    title,
		Subtitle: subtitle,
		Image:    image,
		Buttons:  buttons,
	}
	return Message{
		Content: &Content{
			Attachment: &Attachment{
				Type: "template",
				Payload: &AttachmentPayload{
					Type:     "generic",
					Elements: []*templateElement{&elm},
				},
			},
		},
	}
}

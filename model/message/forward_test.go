package message

import (
	"testing"

	"go_chatible/model/user"
)

func TestBuildForwardMessage(t *testing.T) {
	msg := Message{
		Sender:    &user.User{ID: "1234"},
		Recipient: &user.User{ID: "2345"},
		Timestamp: 1234,
	}
	dummyAttachment := Attachment{Type: "image"}
	dummyAttachment.Payload.URL = "example.com/image"

	textOnlyContent := Content{
		Text: "HelloWorld",
	}
	oneAttachmentContent := Content{
		Attachments: []*Attachment{&dummyAttachment},
	}
	manyAttachmentContent := Content{
		Attachments: []*Attachment{&dummyAttachment, &dummyAttachment},
	}
	textAndOneAttachmentContent := Content{
		Text:        "HelloWorld",
		Attachments: []*Attachment{&dummyAttachment},
	}
	textAndManyAttachmentContent := Content{
		Text:        "HelloWorld",
		Attachments: []*Attachment{&dummyAttachment, &dummyAttachment},
	}

	contents := []Content{textOnlyContent, oneAttachmentContent, manyAttachmentContent, textAndOneAttachmentContent, textAndManyAttachmentContent}
	outputNum := []int{1, 1, 2, 2, 3}
	for i, ctn := range contents {
		msg.Content = &ctn
		msgs := BuildForwardMessage(msg, "4567")
		if len(msgs) != outputNum[i] {
			t.Error("Expect to put out", outputNum[i], "messages but only", len(msgs), "found")
		}
		for _, outputMsg := range msgs {
			if "4567" != outputMsg.Recipient.ID {
				t.Error("Expect output message to have Recipient ID = 4567")
			}

		}
	}
}

package message

type PostbackType string

const (
	Start             PostbackType = "START"
	RequestEnd                     = "REQUEST_END"
	ConfirmEnd                     = "CONFIRM_END"
	ChangePreferrence              = "CHANGE_PREFERRENCE"
	PreferMale                     = "PREFER_MALE"
	PreferFemale                   = "PREFER_FEMALE"
)

func (pl PostbackType) Title() string {
	return map[PostbackType]string{
		"START":              "Start",
		"REQUEST_END":        "End",
		"CONFIRM_END":        "End",
		"CHANGE_PREFERRENCE": "Change gender preferrence",
		"PREFER_MALE":        "Male",
		"PREFER_FEMALE":      "Female",
	}[pl]
}

type Postback struct {
	Title           string `json:"title,omitempty"`
	PostbackPayload string `json:"payload,omitempty"`
}

func MakePostback(tp PostbackType) Postback {
	return Postback{
		Title:           tp.Title(),
		PostbackPayload: string(tp),
	}
}

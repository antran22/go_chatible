package user

type User struct {
	ID          string `json:"id,omitempty" pg:"id,pk"`
	Gender      string `json:"-" pg:"gender,type:gend"`
	Preferrence string `json:"-" pg:"preferrence,type:pref"`
	PartnerID   string `json:"-" pg:"part"`
}

var GenderTypeCommand string = "CREATE TYPE gend AS ENUM ('male', 'female');"
var PrefTypeCommand string = "CREATE TYPE pref AS ENUM ('male', 'female', 'any');"

package user

type User struct {
	ID          string `json:"id,omitempty" sql:"id,pk,unique,type:varchar(20)"`
	Gender      string `json:"gender,omitempty" sql:"gender,type:gend,default:'male'"`
	Preferrence string `json:"-" sql:"preferrence,type:pref,default:'any'"`
	PartnerID   string `json:"-" sql:"partner,type:varchar(20)"`
	Pic         string `json:"profile_pic,omitempty" sql:"pic"`
	FirstName   string `json:"first_name,omitempty" sql:"-"`
	LastName    string `json:"last_name,omitempty" sql:"-"`
	FullName    string `json:"-" sql:"fullname,type:varchar(30)"`
}

func (usr *User) SetFullName() {
	usr.FullName = usr.FirstName + " " + usr.LastName
	if len(usr.FullName) > 30 {
		usr.FullName = usr.FullName[:30]
	}
	usr.FirstName = ""
	usr.LastName = ""
}

var GenderTypeCommand = "CREATE TYPE gend AS ENUM ('male', 'female');"
var PrefTypeCommand = "CREATE TYPE pref AS ENUM ('male', 'female', 'any');"

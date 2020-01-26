package user

import (
	"log"

	"go_chatible/api"
	"go_chatible/model/user"
)

func FetchUser(usr *user.User) error {
	err := api.DB.Select(usr)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			err2 := api.UserFetcher.FetchUserData(usr)
			if err2 != nil {
				return err2
			}
			err2 = api.DB.Insert(usr)
			if err2 != nil {
				log.Printf("unable to save user data with id %s into the database", usr.ID)
			}
		} else {
			return err
		}
	}
	return nil
}

package main

import (
	"log"

	"go_chatible/api"
	"go_chatible/server"

	//userController "go_chatible/controller/user"
	"go_chatible/env"
	//userModel "go_chatible/model/user"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	env.Load("")
	api.InitAPIs()
	server.Serve()
	//usr := userModel.User{ID: "23"}
	//userController.FetchUser(&usr)
	//log.Println(usr)
}

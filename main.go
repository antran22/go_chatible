package main

import (
	"log"

	"go_chatible/api"
	"go_chatible/server"

	"go_chatible/env"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	env.Load("")
	api.InitAPIs()
	defer api.TearDownAPIs()
	server.Serve()
}

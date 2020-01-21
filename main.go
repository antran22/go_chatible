package main

import (
	"go_chatible/env"
	"go_chatible/server"
)

func main() {
	env.Load()
	server.Serve()
}

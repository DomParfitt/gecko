package main

import (
	"github.com/DomParfitt/gecko/server"
)

func main() {
	go server.ServeView("4200")
	server.Serve("8080")
}

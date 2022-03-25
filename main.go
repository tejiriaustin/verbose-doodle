package main

import (
	server "github.com/tejiriaustin/verbose-doodle/internal/rest"
	"log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

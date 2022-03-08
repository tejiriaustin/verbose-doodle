package main

import (
	"fmt"
	server "github.com/tejiriaustin/verbose-doodle/src"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting the on port 8000")
	fmt.Println("Starting go server...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

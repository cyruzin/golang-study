package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", fs)

	log.Println("Listening on port: 8000")
	http.ListenAndServe(":8000", nil)
}

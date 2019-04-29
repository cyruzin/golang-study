package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a Golang server :)"))
	})

	srv := &http.Server{
		Addr:              ":8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	fmt.Println("Listening on port: 8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

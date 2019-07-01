package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(time.Second * 1):
		w.Write([]byte("Message: Context OK!"))
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Error Message: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

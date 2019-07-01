package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	urls := []string{
		"https://www.spotify.com",
		"https://globoesporte.globo.com",
		"https://www.google.com.br",
		"https://www.facebook.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkStatus(url, c)
	}

	for range urls {
		log.Println(<-c)
	}

	close(c)

}

func checkStatus(url string, c chan<- string) {

	resp, err := http.Get(url)

	if err != nil {
		c <- fmt.Sprintf("Error while fetching the url: %s", url)
		c <- fmt.Sprintf("Error: %s", err)
		return
	}

	defer resp.Body.Close()

	u := fmt.Sprintf("\nURL: %s - Status: %s", url, resp.Status)
	c <- u

}

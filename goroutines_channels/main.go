package main

import (
	"fmt"
	"net/http"
)

func main() {

	urls := []string{
		"https://www.google.com.br",
		"https://www.facebook.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkStatus(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

	close(c)

}

func checkStatus(url string, c chan string) {

	resp, err := http.Get(url)

	if err != nil {
		c <- fmt.Sprintf("Error while fetching the url: %s", url)
		c <- fmt.Sprintf("Error: %s", err)
		return
	}

	defer resp.Body.Close()

	c <- resp.Status

}

func printStatus(in <-chan string) {
	for x := range in {
		fmt.Println(x)
	}
}

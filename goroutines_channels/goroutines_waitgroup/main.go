// Waitgropups are userd to wait all the goroutines that are lunched to finish, 
// it notifies that the worker has completed the given task

package main

import (
	"fmt"
	"sync"
	"net/http"
)

func main() {
	urls := []string{
		"https://feelthemovies.com.br/healthcheck",
		"https://www.spotify.com",
		"https://globoesporte.globo.com",
		"https://www.google.com.br",
		"https://www.facebook.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go checkStatus(url, &wg)
	}
	wg.Wait()
}

func checkStatus(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Sprintf("Error: %+v", err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response for %s is %+v \n", url, resp.Status)
	wg.Done()
}

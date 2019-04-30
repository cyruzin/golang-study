package main

import "fmt"

func main() {

	urls := []string{
		"https://www.spotify.com",
		"https://globoesporte.globo.com",
		"https://www.google.com.br",
		"https://www.facebook.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
	}

	// if you don't need a index, use a underscore
	for _, url := range urls {
		fmt.Println(url)
	}

	fmt.Println("-------------------------------")

	for index, url := range urls {
		fmt.Println(index, " - ", url)
	}
}

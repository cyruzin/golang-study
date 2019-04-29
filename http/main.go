package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := httpClient.Get("https://feelthemovies.com.br/v1/recommendations")

	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
	}

	fmt.Printf("%s", data)

}

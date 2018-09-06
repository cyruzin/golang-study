package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	fmt.Println("Creating file...")

	w := []byte("Hello World!\nThat's It!")

	err := ioutil.WriteFile("myfile.txt", w, 0644)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Done!")

	r, err := ioutil.ReadFile("myfile.txt")

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", r)

}

package main

import (
	"log"
	"reflect"
	"strings"
)

// User is a type to demonstrate how reflect works in Golang.
type User struct {
	FirstName string `study:"firstName" json:"first_name"`
	LastName  string `study:"lastName" json:"last_name"`
}

func main() {
	u := User{
		FirstName: "Cyro",
		LastName:  "Dubeux",
	}

	structType := reflect.TypeOf(u)

	checkType(structType, 0)

}

func checkType(t reflect.Type, depth int) {
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			log.Println("Field name: ", f.Name)
			log.Println("Type name: ", f.Type.Name())

			if f.Tag != "" {
				log.Println(strings.Repeat("\t", depth+2), "Tag(s): ", f.Tag)
				log.Println(strings.Repeat("\t", depth+2), "study value: ", f.Tag.Get("study"), "json value: ", f.Tag.Get("json"))
			}
		}
	}
}

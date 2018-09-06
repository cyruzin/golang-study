package main

import (
	"fmt"
)

func main() {

	basket := make(map[interface{}]interface{})
	soccer := make(map[interface{}]interface{})

	basket[0] = "LA Lakers"
	basket[1] = "Houston Rockets"
	basket[2] = "Miami Heat"
	basket[3] = "Orlando Magic"
	basket[4] = "Washington Wizards"

	soccer["B"] = "Barcelona"
	soccer["RM"] = "Real Madrid"
	soccer["J"] = "Juventus"
	soccer["P"] = "PSG"
	soccer["BM"] = "Bayern Munich"

	listMap(basket)
	fmt.Println("")
	listMap(soccer)

}

func listMap(m map[interface{}]interface{}) {
	for i, v := range m {
		fmt.Println(i, " - ", v)
	}
}

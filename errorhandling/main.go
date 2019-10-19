package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	number1 := 25.0
	number2 := -1.0

	result1, err := Math(number1, "sqrt")
	if err != nil {
		fmt.Println("math error: ", err)
	} else {
		fmt.Printf("The square root of %f is %f\n", number1, result1)
	}

	result2, err := Math(number2, "sqrt")
	if err != nil {
		fmt.Println("math error: ", err)
	} else {
		fmt.Printf("The square root of %f is %f", number2, result2)
	}

}

func Math(number1 float64, calc string) (float64, error) {
	const errMessage = "mathfunction: "
	switch calc {
	case "sqrt":
		result, err := Sqrt(number1)
		if err != nil {
			errMsg := errMessage + err.Error()
			return 0, errors.New(errMsg)
		}
		return result, nil
	default:
		return 0, errors.New(errMessage + "unkown type of calc, use 'sqrt'")
	}

}

func Sqrt(number float64) (float64, error) {
	const errMessage = "square root of negative number"
	if number < 0 {
		return 0, errors.New(errMessage)
	}
	return math.Sqrt(number), nil
}

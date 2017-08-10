package main

import (
	"fmt"
	"math"
)

func main() {
    for i := 0; i < 100; i++ {
	multipleOfThree := math.Mod(float64(i), 3.0 )  == 0
	multipleOfFive := math.Mod(float64(i), 5.0 )  == 0

	if multipleOfThree && multipleOfFive {
	    fmt.Println("fizzbuzz")
	} else if multipleOfThree {
	    fmt.Println("fizz")
	} else if multipleOfFive {
	    fmt.Println("buzz")
	} else {
	    fmt.Println(i)
	}
    }
}

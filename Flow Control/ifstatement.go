package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//highlighting the use of if else if and else statements
func powpow(x, y, z float64) {
	if x == 0 {
		fmt.Println("Does this have any meaning?")
	} else if x > 0 {
		fmt.Println("Does anything have any meaning?")
	} else if x < 0 {
		fmt.Println("Do we create our own meaning?")
	} else {
		fmt.Println("Maybe this has more meaning than what we give it")
	}
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

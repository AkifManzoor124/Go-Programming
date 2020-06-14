package main

import "fmt"

func runtime(location string) {
	fmt.Println("Where would you like to run?")
	switch location {
	case "Milton":
		fmt.Println("1:00:00")
	case "Toronto":
		fmt.Println("2:50:23")
	default:
		fmt.Println("Cannot find run time for location")
	}
}

//We can write a switch with no condition to write clean long chains of if-then-else statements
func adviceOnShoes(shoeType string) {
	switch {
	case shoeType == "Slippers":
		fmt.Println("Goodluck")
	case shoeType == "Running Shoes":
		fmt.Println("Good Choice")
	default:
		fmt.Println("Oooooo Boy")
	}
}

func main() {
	fmt.Println("Running Switch Example")
	runtime("Toronto")
	adviceOnShoes("None")
}

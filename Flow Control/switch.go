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

func main() {
	fmt.Println("Running Switch Example")
	runtime("Toronto")
}

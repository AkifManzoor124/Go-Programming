package main

import "fmt"

type Car struct {
	name string
	make string
}

func main() {
	teslabuild := Car{"Tesla", "Model S"}
	fmt.Println(teslabuild)
	fmt.Println("The make of the car is:", teslabuild.name)

	//you can access the struct fields using a pointer
	fordbuild := Car{"Ford", "fordie"}
	p := &fordbuild

	fmt.Println(p.name)
}

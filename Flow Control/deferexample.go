package main

import "fmt"

//defer statements are put onto a stack and come out of the stack with the first in last out methodology
func main() {
	defer fmt.Println("You ain't shit")

	fmt.Println("Get me a sandwich")

	//1 will go in the stack first and will be printed last
	//9 will print first
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

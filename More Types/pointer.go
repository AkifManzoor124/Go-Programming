package main

import "fmt"

func main() {
	//pointers hold the memory address to a value

	i, j := 42, 2701

	//* is used to find the value
	//& is used to point to the assigned variable
	//p is a pointer to i
	p := &i
	//* will find the value of i using the pointer we defined for it
	fmt.Println(*p)
	//We can set i to another value using * which gets us the value of i
	*p = 21
	fmt.Println(i)

	p = &j       //assigning p to the pointer of j
	*p = *p / 37 //value of j is set to j / 37
	fmt.Println(j)

}

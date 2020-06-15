package main

import "fmt"

func main() {
	//arrays are defined by the variable name, [# of elements] and the type of elements the array can hold
	var goals [2]string
	goals[0] = "Graduate Computer Engineering"
	goals[1] = "Wake up at 10am"

	fmt.Println(goals)
	fmt.Println(goals[0])
	fmt.Println(goals[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slices are dynamically-sized and flexible views into the array
	var s []int = primes[1:4]
	fmt.Println(s)
	s[0] = 23

	fmt.Println(primes)

	//We can find the capacity of the slice using cap(slice) which gives us the # of elements the slice can hold
	//counting from the first element in the slice
	//This is not the lenght of the array the slice is referencing
	fmt.Println(cap(s))

	//We can find the number of items within the slice using len(s)
	fmt.Println(len(s))

	s = primes[1:6]

	fmt.Println(s)

	//To create a dynamically sized array, we can use the make([]type, lenght) to create a slice which references a zeroed array

	//this will create an array with 6 0's
	a := make([]string, 6)
	//this will create an empty array with a capacity of 6 values
	b := make([]string, 0, 6)

	fmt.Println(a)
	fmt.Println(b)

	//We can append to a slice using the append keyword
	var c []int
	c = append(c, 0,1,2,3,3,3,4,5,2,3,2)
	fmt.Println(c)

	//we can use the range keyword which iterates over a slice
	for i := range c{
		fmt.Printf("%d",c[i])
	}
}

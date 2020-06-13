package main

import "fmt"

func main(){
	//There is no parantheses for the for loop 
	for i := 0; i < 10; i++ {
		fmt.Println(i)	
	}

	//Go uses for as the for loop and as the while loop
	//We need to increase j or else, we will only print 0s
	var j int
	for j < 1000 {
		j = j + 1 
		fmt.Println(j)
	}
}

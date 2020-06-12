package main

import(
	"fmt"
	"math"
)

//we can define variables in a block just like imports
var{
	foo, bar = 1, math.Pi
	x,y bool = true, false
	i,j int = 2, 1
}

func main(){
	fmt.Printf("Type: %T Value %v\n", foo,bar)
	fmt.Printf("Type: %T Value %v\n", x,y)
	fmt.Printf("Type: %T Value %v\n", i,j)
}
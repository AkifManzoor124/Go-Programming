package main

import(
	"fmt"
	"math"
)

//we can define variables in a block just like imports
//this eliminates the need to put var infront of each and every varialbe definition
var(
	foo, bar = 1, math.Pi
	x,y bool = true, false
	i,j int = 2, 1

	//zero values example
	jk int
	balalala bool
	s string
)

func main(){
	fmt.Printf("Type: %T Value %v\n", foo,foo)
	fmt.Printf("Type: %T Value %v\n", x,x)
	fmt.Printf("Type: %T Value %v\n", i,i)

	fmt.Printf("Can't touch this %v %v %v",jk,balalala,s)
}
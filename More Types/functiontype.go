package main

import(
	"fmt"
	"math"
)

//Functions are variables in go and can be used as such
//They can be passed to other functions
//function inception?

//fn is the name of the function within the function definition
func compute(fn func(float64,float64) float64) float64{
	return fn(3,4)
}

func main(){
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(math.Pow(3,4))
}
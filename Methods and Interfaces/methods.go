package main

import "fmt"

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Go does not have classes but we can define methods for structs
func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + x.Y*v.Y)
}

//you can also declare methods on non-struct types too
type Myvar int

func (f Myvar) Abs() float64{
	if f < 0{
		return f*2
	}
	return f*3
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
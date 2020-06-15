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

func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
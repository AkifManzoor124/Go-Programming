package main

import "fmt"

var foo, bar bool //<-- Global Variables 
//boolean values are set to false if not initialized

var x,y,z bool = false, false, true
//x,y,z must be defined as booleans now

func main(){
	var i int
	var j,k,l = 1, 2, true
	//freedom to initialize j,k,l as what you want
	fmt.Println(i, foo, bar)
}
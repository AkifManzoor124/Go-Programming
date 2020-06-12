package main

import "fmt"

func double(x,y int) (a,b int){
	a = x * 2
	b = y * 2
	return //<== This is called a "Naked" return statement
}

func main(){
	fmt.Println(double(2,4))
}
package main

import "fmt"

const (
	a=42
	b=42.44
	c="james bond"
)

func main (){
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)

}
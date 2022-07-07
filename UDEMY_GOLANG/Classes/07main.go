package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main (){
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)

}
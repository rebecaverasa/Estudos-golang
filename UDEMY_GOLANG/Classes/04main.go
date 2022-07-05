package main

import "fmt"

var a int

type caneta int

var b caneta

func main() {
	a=42
	b=43
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)

	a=int(b) //convertion b de caneta pra int 

	fmt.Printf("%v\t%T\n", a, a)

}

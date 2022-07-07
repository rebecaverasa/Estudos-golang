package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10) //1 iota
	mb = 1 << (iota * 10) //2 iota (10*2=20) ou seja, 20 casas pra frente
	gb = 1 << (iota * 10) //3 iota (10*3=30) ou seja, 30 casas pra frente
)

func main() {
	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}

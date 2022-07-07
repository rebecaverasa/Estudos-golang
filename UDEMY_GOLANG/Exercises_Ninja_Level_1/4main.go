package main

import "fmt"

type borracha int

var x borracha

func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x=42
	fmt.Println(x)

}
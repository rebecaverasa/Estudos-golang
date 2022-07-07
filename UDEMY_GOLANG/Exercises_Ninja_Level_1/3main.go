package main

import "fmt"

var x = 42
var y = "james bond"
var z = true

func main (){
	
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Printf("%T",s)
}
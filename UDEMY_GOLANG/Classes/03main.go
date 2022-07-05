package main

import "fmt"

var y = 42

func main (){
	fmt.Println(y)
	fmt.Printf("%T \n", y)
	fmt.Printf("%b \n", y)
	fmt.Printf("%x \n", y)
	fmt.Printf("%#x \n", y)
	
	y=911

	fmt.Printf("%v\t%T\t%b\t%x\t%#x", y, y, y, y, y)

}
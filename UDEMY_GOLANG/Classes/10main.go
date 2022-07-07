package main

import "fmt"

func main (){
	//for init; condition; post(após)
	for i:=0; i<=10; i++{
		fmt.Printf("The outer loop: %v.\n", i)
		for j:=0; j<3; j++{
			fmt.Printf("\tThe inner loop: %v\n", j)
		}
	}
}
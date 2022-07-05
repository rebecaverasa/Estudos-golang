package main

import "fmt"

	func main(){
		area := func(width, height float64)float64{
			return width * height
		}
		fmt.Println(area(2,3))	

	}
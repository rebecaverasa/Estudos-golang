package main

import (
	"fmt"
)

func rectangleArea(width, height float64)(result float64, result2 float64){
result = width*height
return
}

func main(){
	res1, res2 := rectangleArea(2,3)
	fmt.Println("rectangleArea =", res1, res2)
}
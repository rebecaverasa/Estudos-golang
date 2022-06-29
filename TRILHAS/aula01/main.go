//resolver um numero fatorial de forma iterativa

package main

import (
	"fmt"
)

func fatorial (numero int) int {
	result:=1
	for i := 1; i<=numero; i++{
		result *=i
	}
	return result
}

func main (){
	fmt.Println(fatorial (10))
}
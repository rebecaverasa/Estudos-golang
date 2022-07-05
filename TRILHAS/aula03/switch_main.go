package main

import "fmt"
import "time"

func main() {
	value := 10
	switch value {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("default")
	}
	horaAgora()
}

func horaAgora(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("bom dia")
	case t.Hour() < 18:
		fmt.Println("boa tarde")
	default:
		fmt.Println("boa noite")
	}
}
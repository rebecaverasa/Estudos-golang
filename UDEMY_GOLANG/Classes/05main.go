package main

import "fmt"

func main (){
	s:="H"
	fmt.Println(s)
	
	bs:=[]byte(s) //Um slide da variavel s, para saber quantos bytes tem a variavel s
	fmt.Println(bs)

	n:= bs[0] 
	fmt.Println(n)
	fmt.Printf("%T\n", n) //saber o tipo de n, decimal
	fmt.Printf("%b\n", n) //saber o binario
	fmt.Printf("%X\n", n) //saber o hexadecimal



}
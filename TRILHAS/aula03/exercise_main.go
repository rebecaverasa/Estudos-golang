package main

import (
	"fmt"
)

type key interface {
	Press() interface {}
}

type keyboard struct {
	LastKeyPresses interface {}
	Keys []Key 
}

type KeyHelloWorld struct {}

func (Khw KeyHelloWorld) Press() interface {} {
	fmt.Println("Hello World")
	return "Hello World"
}

func main (){
	keyboard := keyboard{Keys: []Keys{}}
}
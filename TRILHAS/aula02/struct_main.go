package main

import "fmt"

type User struct{
	name string
	age int
}

func main(){
	user1:= User{name: "rebeca", age:26}
	fmt.Println(user1)
}
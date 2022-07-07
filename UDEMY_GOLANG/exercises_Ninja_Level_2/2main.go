package main

import "fmt"

func main(){
	a:= 42==3
	b:= 42<=3
	c:= 42>=3
	d:= 42!=3
	e:= 42<3
	f:= 42>3

	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)
	fmt.Printf("%v\t%T\n", e, e)
	fmt.Printf("%v\t%T\n", f, f)

}
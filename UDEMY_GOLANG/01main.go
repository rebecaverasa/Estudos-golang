package main
import "fmt"
func main (){
	fmt.Println("Ola")
	teste()

	for i:=0; i<=100; i++{
		if i%2==0{
			fmt.Print(i)
		}
	}
	fim()
}

func teste(){
	fmt.Println("mudança")
}

func fim(){
	fmt.Println("\nAqui é o fim")
}
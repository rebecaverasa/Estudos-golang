package main

import "fmt"

var y = 43 //declarar com var é pra ser usado fora da funcao main.

var z int

//declaramos que ha uma variavel Z que é do tipo inteiro. Qnd nao temos o valor, precisa-se dizer o tipo.
//enquanto nao ha valor atribuido para uma variavel: 0 inteiros, false boolean, 0.0 float, -- string.

func main() {
	x := 42 //declaracao curta é pra ser usada dentro da funcao main.

	fmt.Println(x)

	fmt.Println(y)

	fmt.Println("z")

	foo()
}

func foo() {
	fmt.Println("again:", y)
	//no caso, como a var y esta fora da func main, ela pode ser usada em todo o codigo.
	// se fosse printar a x nao daria certo, pq ela esta limitada a funcao main.
}

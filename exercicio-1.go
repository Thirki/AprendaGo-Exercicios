package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	// Desafio 1
	fmt.Println(x, y, z)
	// Desafio 2
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// Desafio bonus
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}

// - Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
//     1. 42
//     2. "James Bond"
//     3. true
// - Agora demonstre os valores nestas variáveis, com:
//     1. Uma única declaração print
//     2. Múltiplas declarações print

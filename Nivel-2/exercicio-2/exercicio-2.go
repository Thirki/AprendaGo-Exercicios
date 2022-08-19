package main

import "fmt"

const (
	a = iota
	b
)

func main() {
	fmt.Printf("%v == %v \t %v", a, b, (a == b))
	fmt.Printf("\n%v != %v \t %v", a, b, (a != b))
	fmt.Printf("\n%v <= %v \t %v", a, b, (a <= b))
	fmt.Printf("\n%v < %v \t %v", a, b, (a < b))
	fmt.Printf("\n%v >= %v \t %v", a, b, (a >= b))
	fmt.Printf("\n%v > %v \t %v", a, b, (a > b))
}

// - Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
// 	- ==
// 	- !=
// 	- <=
// 	- >=
// 	- <
// 	- >
// - Demonstre estes valores.

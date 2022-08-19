package main

import "fmt"

const (
	_ = iota + 2022
	ano1
	ano2
	ano3
	ano4
)

func main() {
	fmt.Printf("Primeiro Ano\t- \t%v\n", ano1)
	fmt.Printf("Segundo Ano\t- \t%v\n", ano2)
	fmt.Printf("Terceiro Ano\t- \t%v\n", ano3)
	fmt.Printf("Quarto Ano\t- \t%v\n", ano4)
}

// @TODO
// - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.

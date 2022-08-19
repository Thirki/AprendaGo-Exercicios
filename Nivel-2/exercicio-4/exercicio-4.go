package main

import "fmt"

var x int = 10

func main() {
	fmt.Printf("Variavel inicial\nTipo: %T\tDecimal: %d\tBinario: %b\tHexadecimal: %#x", x, x, x, x)
	y := x << 1
	fmt.Printf("\nVariavel deslocada\nTipo: %T\tDecimal: %d\tBinario: %b\tHexadecimal: %#x", y, y, y, y)
}

// TODO:
// - Crie um programa que:
//     - Atribua um valor int a uma variável
//     - Demonstre este valor em decimal, binário e hexadecimal
//     - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
//     - Demonstre esta outra variável em decimal, binário e hexadecimal

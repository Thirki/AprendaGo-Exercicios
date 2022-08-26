package main

import "fmt"

func main() {
	numberList := []int{1, 2, 50, 33, 10}
	FirstSumOfNumbers := somaValores(numberList...)
	SecondSumOfNumbers := secondSomaValores(numberList)
	fmt.Printf("A soma de %v\n é de %v ", numberList, FirstSumOfNumbers)
	fmt.Printf("\nA soma de %v\n é de %v ", numberList, SecondSumOfNumbers)
}

func somaValores(numberList ...int) int {
	total := 0
	for _, number := range numberList {
		total += number
	}
	return total
}

func secondSomaValores(numberList []int) int {
	total := 0
	for _, number := range numberList {
		total += number
	}
	return total
}

// // @TODO:
// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

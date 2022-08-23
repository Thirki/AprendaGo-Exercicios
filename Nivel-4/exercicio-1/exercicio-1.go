package main

import "fmt"

func main() {
	nums := [5]int{6, 7, 8, 9, 10}

	for i, num := range nums {
		println(i, "-", num)
	}

	fmt.Printf("Tipo: %T", nums)
}

// @TODO:
// - Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.

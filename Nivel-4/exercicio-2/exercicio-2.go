package main

import "fmt"

func main() {
	numbers := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for _, number := range numbers {
		fmt.Println(number)
	}

	fmt.Printf("%T, %v", numbers, numbers)
}

// @TODO:
// - Usando uma literal composta:
//     - Crie uma slice de tipo int
//     - Atribua 10 valores a ela
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.

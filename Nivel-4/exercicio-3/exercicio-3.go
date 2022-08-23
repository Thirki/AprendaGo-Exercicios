package main

import "fmt"

func main() {
	numbers := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(numbers[:3])
	fmt.Println(numbers[4:])
	fmt.Println(numbers[1:7])
	fmt.Println(numbers[2 : len(numbers)-1])

	fmt.Printf("%T, %v", numbers, numbers)
}

// @TODO:
// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
//     - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//     - Do quinto ao último item do slice (incluindo o último item!)
//     - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//     - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//     - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

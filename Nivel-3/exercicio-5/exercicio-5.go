package main

import "fmt"

func main() {
	for index := 10; index <= 100; index++ {
		fmt.Printf("%v %% 4 = %v\n", index, index%4)
	}
}

// @TODO:
// Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

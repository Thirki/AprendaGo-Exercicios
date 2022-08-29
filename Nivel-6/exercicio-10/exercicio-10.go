package main

import "fmt"

func main() {
	contadorA := contaExecucoes()
	fmt.Println(contadorA())
	fmt.Println(contadorA())
	fmt.Println(contadorA())
	contadorB := contaExecucoes()
	fmt.Println(contadorB())
	fmt.Println(contadorB())
	fmt.Println(contadorB())
}

func contaExecucoes() func() int {
	execucoes := 0
	return func() int {
		execucoes += 1
		return execucoes
	}
}

// @TODO:
// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

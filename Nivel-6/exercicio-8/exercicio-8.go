package main

import "fmt"

func main() {
	callback := getCallback()
	callback()
}

func getCallback() func() {
	return func() {
		fmt.Println("que role, heim")
	}
}

// @TODO:
// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.

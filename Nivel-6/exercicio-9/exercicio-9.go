package main

import "fmt"

func main() {
	componente(setState)
}

func setState() {
	fmt.Println("Setei o estado para X")
}

func componente(stateFunction func()) {
	fmt.Println("Que sdds do front...")
	stateFunction()
}

// @TODO:
// - Callback: passe uma função como argumento a outra função.

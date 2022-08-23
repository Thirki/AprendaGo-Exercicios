package main

import "fmt"

func main() {
	idade := 15
	switch {
	case idade == 18:
		fmt.Println("AGORA PODE BEBER!")
	case idade < 18:
		fmt.Println("Não pode beber")
	case idade > 18:
		fmt.Println("Pode beber")
	default:
		fmt.Println("idade não identificada")
	}
}

// @TODO:
// - Crie um programa que utilize a declaração switch, sem switch statement especificado.

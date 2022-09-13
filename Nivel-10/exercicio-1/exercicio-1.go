package main

import (
	"fmt"
)

// Usando uma função anônima auto-executável
func main() {
	primeiraForma()
	segundaForma()
}

func primeiraForma() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func segundaForma() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

// @TODO:
// - Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
// - Usando uma função anônima auto-executável
// - Usando buffer

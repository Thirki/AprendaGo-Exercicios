package main

import "fmt"

func main() {
	c1 := make(chan int)

	go func(channel chan<- int) {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}(c1)

	for value := range c1 {
		fmt.Println(value)
	}
}

// @TODO:
// - Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

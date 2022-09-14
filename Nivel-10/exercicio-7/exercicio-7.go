package main

import "fmt"

func main() {
	c1 := make(chan int)

	for i := 0; i < 10; i++ {
		go criar_canal(10, c1)
	}

	for x := 0; x < 100; x++ {
		fmt.Println(x, "\t", <-c1)
	}
}

func criar_canal(loop_times int, c1 chan<- int) {
	for i := 0; i < loop_times; i++ {
		c1 <- i
	}
}

// @TODO:
// - Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()

	return c
}

func receive(c1 <-chan int, c2 chan int) {
	fmt.Println(c2)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case <-c2:
			return
		}
	}
}

// @TODO:
// - Utilizando este código: https://play.golang.org/p/MvL6uamrJP
// - ...use um select statement para demonstrar os valores do canal.

package main

import "fmt"

func main() {
	numero := 20
	if numero == 10 {
		fmt.Println("o numero é 10!")
	} else if numero < 10 {
		fmt.Println("o numero é menor que 10!")
	} else {
		fmt.Println("o numero é maior que 10!")
	}
}

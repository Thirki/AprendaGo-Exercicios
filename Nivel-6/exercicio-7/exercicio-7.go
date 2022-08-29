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
// - Atribua uma função a uma variável.
// - Chame essa função.

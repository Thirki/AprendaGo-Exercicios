package main

import "fmt"

func main() {
	func(prova string) {
		fmt.Println("ISSO FOI EXECUTADO AUTOMAGICAMENTE E EU TENHO A PROVA")
		fmt.Printf("Prova: %v", prova)
	}("EU SOU A PROVA")
}

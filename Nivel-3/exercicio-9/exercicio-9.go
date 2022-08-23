package main

import "fmt"

var esporteFavorito string

func main() {
	esporteFavorito = "Basquete"
	switch esporteFavorito {
	case "Basquete":
		fmt.Println("Esse é bom de cesta!")
	case "Futebol":
		fmt.Println("Brasileiro nato")
	case "Golf":
		fmt.Println("Alguem com menos de 100 anos joga isso ??!!")
	default:
		fmt.Println("Esporte não identificado")
	}
}

// @TODO:
// - Crie um programa que utilize a declaração switch
// 	onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {
	fmt.Println(nivelDeGo(5))
	numberLevel, stringLevel := megaNivelDeGo(10)
	fmt.Printf("Seu nivel de Go é de %v, pelos meus calculos você é um ninja go nivel %v", numberLevel, stringLevel)

}

func nivelDeGo(lvl int) int {
	return lvl
}

func megaNivelDeGo(lvl int) (int, string) {
	goLevel := "undefined"
	switch {
	case lvl >= 0 && lvl <= 5:
		goLevel = "Baixo"
	case lvl > 5 && lvl <= 7:
		goLevel = "Médio"
	case lvl > 7:
		goLevel = "MESTRE NINJA!"
	}

	return lvl, goLevel
}

// @TODO:
// - Exercício:
//     - Crie uma função que retorne um int
//     - Crie outra função que retorne um int e uma string
//     - Chame as duas funções
//     - Demonstre seus resultados

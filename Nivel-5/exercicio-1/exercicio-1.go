package main

import "fmt"

type IPessoa = struct {
	Nome                         string
	Sobrenome                    string
	Sabores_favoritos_de_sorvete []string
}

func main() {
	pessoa1 := IPessoa{
		Nome:                         "Henrique",
		Sobrenome:                    "Gomes",
		Sabores_favoritos_de_sorvete: []string{"Creme", "Milho"},
	}

	pessoa2 := IPessoa{"Francine", "Milanes", []string{"Milho", "Chocomenta"}}

	fmt.Println(pessoa1.Nome)

	for _, sabores := range pessoa1.Sabores_favoritos_de_sorvete {
		fmt.Println("\t", sabores)
	}

	fmt.Println(pessoa2.Nome)

	for _, sabores := range pessoa2.Sabores_favoritos_de_sorvete {
		fmt.Println("\t", sabores)
	}
}

// @TODO:
// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

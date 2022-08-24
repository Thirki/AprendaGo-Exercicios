package main

import "fmt"

type IPessoa = struct {
	Nome                         string
	Sobrenome                    string
	Sabores_favoritos_de_sorvete []string
}

func main() {
	pessoas := make(map[string]IPessoa)
	pessoas["Gomes"] = IPessoa{
		Nome:                         "Henrique",
		Sobrenome:                    "Gomes",
		Sabores_favoritos_de_sorvete: []string{"Creme", "Milho"},
	}

	pessoas["Milanes"] = IPessoa{"Francine", "Milanes", []string{"Milho", "Chocomenta"}}

	for key, pessoa := range pessoas {
		fmt.Println(key)
		fmt.Println("\t", pessoa)
	}
}

// @TODO:
// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

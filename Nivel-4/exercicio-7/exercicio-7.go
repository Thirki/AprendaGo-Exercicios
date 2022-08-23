package main

import "fmt"

func main() {
	gruposDeAlunos := [][]string{
		{"Henrique", "Gomes", "Jogar"},
		{"Renato", "Gomes", "Estudar"},
		{"Francine", "Milanes", "Ler"},
	}

	for _, alunos := range gruposDeAlunos {
		fmt.Println(alunos)
		for _, informacoes := range alunos {
			fmt.Println(informacoes)
		}
	}
}

// @TODO:
// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

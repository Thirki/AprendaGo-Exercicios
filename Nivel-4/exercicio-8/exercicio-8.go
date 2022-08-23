package main

import "fmt"

func main() {
	users := map[string][]string{
		"Henrique_Gomes":   {"Jogar Valorant", "ler manga"},
		"Francine_Milanes": {"Ler livros", "jogar the sims"},
	}
	fmt.Println(users)
	for key, values := range users {
		fmt.Println(key)
		for index, value := range values {
			fmt.Println("\t", index, value)
		}
	}
}

// @TODO:
// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

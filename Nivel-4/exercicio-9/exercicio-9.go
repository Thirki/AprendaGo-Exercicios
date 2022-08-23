package main

import "fmt"

func main() {
	users := map[string][]string{
		"Henrique_Gomes":   {"Jogar Valorant", "ler manga"},
		"Francine_Milanes": {"Ler livros", "jogar the sims"},
	}
	fmt.Println(users)
	users["Reyna"] = []string{"Trabalhar na Riot", "Comprar roupas estilosas"}
	fmt.Println(users)
	for key, values := range users {
		fmt.Println(key)
		for index, value := range values {
			fmt.Println("\t", index, value)
		}
	}
}

// @TODO:
// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

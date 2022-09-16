package main

import (
	"fmt"
	"v2/Nivel-12/exercicio-1/cachorro"
)

func main() {
	fmt.Println(cachorro.Idade(10))
}

// @TODO:
// - Crie um package "cachorro".
//     - Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
//     - Documente seu código com comentários, e utilize a função Idade na sua função main.
// - Rode seu programa para verificar se ele funciona.
// - Rode um local server com godoc e leia sua documentação.

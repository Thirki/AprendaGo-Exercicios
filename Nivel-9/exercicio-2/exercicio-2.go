package main

import "fmt"

type pessoa struct {
	nome  string
	frase string
}

type humano interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Printf("%v disse: %v", p.nome, p.frase)
}

func dizerAlgumaCoisa(h humano) {
	h.falar()
}

func main() {
	user := pessoa{"Henrique", "to com fome"}
	dizerAlgumaCoisa(&user)
}

// @TODO:
// - Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//     - Crie um tipo para um struct chamado "pessoa"
//     - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//     - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//     - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//     - Demonstre no seu código:
//         - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//         - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"

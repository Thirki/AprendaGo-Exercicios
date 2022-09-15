package main

import "fmt"

type erroEspecial struct {
	nome string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("É %v, parece que deu erro", e.nome)
}

func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(erroEspecial).nome, "\nno método Error, eu tenho:", e)
}

func main() {
	arg := erroEspecial{"Henrique"}
	erroComoParametro(arg)
}

// @TODO:
// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
// - Crie uma função que tenha um valor do tipo error como parâmetro.
// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.

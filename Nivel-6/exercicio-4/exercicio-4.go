package main

import (
	"fmt"
	"strconv"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	user := pessoa{"Henrique", "Gomes", 23}
	greethings := pessoa.retornaTudo(user)
	fmt.Println(greethings)
}

func (p pessoa) retornaTudo() string {
	return "Meu nome é " + p.nome + " " + p.sobrenome + " tenho " + strconv.Itoa(p.idade) + " anos"
}

// @TODO:
// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

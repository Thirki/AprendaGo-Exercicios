package main

import "fmt"

type Ix int

var x Ix

func main() {
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)
	x = 42
	fmt.Println("Novo Valor:", x)
}

// Crie um tipo. O tipo subjacente deve ser int.
// - Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// - Na função main:
//     1. Demonstre o valor da variável "x"
//     2. Demonstre o tipo da variável "x"
//     3. Atribua 42 à variável "x" utilizando o operador "="
//     4. Demonstre o valor da variável "x"

package main

import "testing"

func IsOdd(number int) bool {
	return (number % 2) != 0
}

func VerificaSeÉPar(t *testing.T) {
	v := IsOdd(5)
	if !v {
		t.Error("é par ", v)
	}
}

// - Nos capítulos seguintes, uma das coisas que veremos é testes.
// - Para testar sua habilidade de se virar por conta própria... desafio:
//     - Utilizando as seguintes fontes: https://godoc.org/testing & http://www.golang-book.com/books/intr...
//     - Tente descobrir por conta própria como funcionam testes em Go.
//     - Pode usar tradutor automático, pode rodar código na sua máquina, pode procurar no Google. Vale tudo.
//     - O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.

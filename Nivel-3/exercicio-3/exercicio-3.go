package main

import (
	"fmt"
	"time"
)

func main() {
	currentYear := time.Now().Year()
	anoAtual := 1998
	for anoAtual <= currentYear {
		fmt.Println(anoAtual)
		anoAtual++
	}
}

// @TODO:
// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

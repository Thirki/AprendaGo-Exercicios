package main

import (
	"fmt"
	"time"
)

func main() {
	yearLimit := time.Now().Year()
	bornYear := 1998
	currentYear := bornYear
	for {
		if currentYear > yearLimit {
			break
		}
		fmt.Println(currentYear)
		currentYear++
	}
}

// @TODO:
// Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {
	for unicodeNumber := 65; unicodeNumber <= 90; unicodeNumber++ {
		fmt.Printf("%v:\n", unicodeNumber)
		for repeatTimes := 1; repeatTimes <= 3; repeatTimes++ {
			fmt.Printf("\t%v:, %c\n", repeatTimes, unicodeNumber)
		}
	}
}

// @TODO:
// - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
// - Por exemplo:
// 65
// U+0041 'A'
// U+0041 'A'
// U+0041 'A'
// 66
// U+0042 'B'
// U+0042 'B'
//     // U+0042 'B'
// ...e por aí vai.

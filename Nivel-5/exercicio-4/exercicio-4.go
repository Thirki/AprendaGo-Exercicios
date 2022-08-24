package main

import "fmt"

func main() {
	IComputador := struct {
		perifericos []string
		marca       map[string]string
	}{
		perifericos: []string{"Mouse", "Teclado", "Webcam"},
		marca:       map[string]string{"Ferrugem": "Na pintura toda"},
	}

	fmt.Println(IComputador)
}

// @TODO:
// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

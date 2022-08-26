package main

import (
	"fmt"
	"math"
)

var pi float64 = math.Pi

type IQuadrado struct {
	lado float64
}

type ICirculo struct {
	raio float64
}

// Aqui tipa como figura tudo que tem o campo Area()
type figura interface {
	Area() float64
}

func (q IQuadrado) Area() float64 {
	return 4 * q.lado
}

func (c ICirculo) Area() float64 {
	return 2 * c.raio * pi
}

// Aqui adiciona a função info apenas para os tipos figuras
func info(f figura) float64 {
	switch f.(type) {
	case ICirculo:
		fmt.Println("Agora eu to no Circulo")
	case IQuadrado:
		fmt.Println("Agora eu to no Quadrado")
	}
	return f.Area()
}

func main() {
	meuQuadradoLindo := IQuadrado{lado: 10}
	meuCirculoLindo := ICirculo{raio: 2}

	// Aqui usa o metodo dentro de figuras para dois tipos diferentes que herdaram graças ao metodo Area que os dois tipos tem
	infoQuadrado := info(meuQuadradoLindo)
	infoCirculo := info(meuCirculoLindo)

	fmt.Println(infoQuadrado)
	fmt.Println(infoCirculo)
}

// @TODO:
// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

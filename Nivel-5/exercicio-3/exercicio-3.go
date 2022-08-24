package main

import "fmt"

type IVeiculo struct {
	Portas int
	Cor    string
}

type ICaminhonete struct {
	IVeiculo
	traçãoNasQuatro bool
}

type ISedan struct {
	IVeiculo
	modeloLuxo bool
}

func main() {
	caminhonete := ICaminhonete{IVeiculo{4, "Preta"}, true}
	carroDeSenior := ISedan{IVeiculo{2, "Vermelho"}, true}
	fmt.Printf("Type: %T\tValue: %v\n", caminhonete, caminhonete)
	fmt.Printf("Type: %T\tValue: %v\n", carroDeSenior, carroDeSenior)
	fmt.Println(caminhonete.Cor)

	fmt.Println(carroDeSenior.Cor)
}

// @TODO:
// - Crie um novo tipo: veículo
//     - O tipo subjacente deve ser struct
//     - Deve conter os campos: portas, cor
// - Crie dois novos tipos: caminhonete e sedan
//     - Os tipos subjacentes devem ser struct
//     - Ambos devem conter "veículo" como struct embutido
//     - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
//     - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
// - Usando os structs veículo, caminhonete e sedan:
//     - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//     - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// - Demonstre estes valores.
// - Demonstre um único campo de cada um dos dois.

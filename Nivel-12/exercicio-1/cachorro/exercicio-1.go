// Package cachorro tras otima funcionalidades ao "AUsuario"
package cachorro

// Idade é responsavel por transformar a idade humana em idade de cachorro
func Idade(age int) int {
	multiplicador := 7
	return age * multiplicador
}

// @TODO:
// - Crie um package "cachorro".
//     - Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
//     - Documente seu código com comentários, e utilize a função Idade na sua função main.
// - Rode seu programa para verificar se ele funciona.
// - Rode um local server com godoc e leia sua documentação.

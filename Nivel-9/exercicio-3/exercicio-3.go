package main

import (
	"fmt"
	"runtime"
	"sync"
)

var valorDaContagem int = 0
var wg sync.WaitGroup

func main() {
	loopTimes := 100
	wg.Add(loopTimes)
	for i := 0; i < loopTimes; i++ {
		go func() {
			defer wg.Done()
			v := valorDaContagem
			runtime.Gosched()
			v++
			valorDaContagem = v
		}()
	}
	wg.Wait()
	fmt.Println("valor final:", valorDaContagem)
}

// @TODO:
// - Utilizando goroutines, crie um programa incrementador:
//     - Tenha uma variável com o valor da contagem
//     - Crie um monte de goroutines, onde cada uma deve:
//         - Ler o valor do contador
//         - Salvar este valor
//         - Fazer yield da thread com runtime.Gosched()
//         - Incrementar o valor salvo
//         - Copiar o novo valor para a variável inicial
//     - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//     - Demonstre que há uma condição de corrida utilizando a flag -race

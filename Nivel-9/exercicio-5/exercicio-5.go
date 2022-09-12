package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var valorDaContagem int64 = 0
var wg sync.WaitGroup

func main() {
	loopTimes := 10
	wg.Add(loopTimes)
	for i := 0; i < loopTimes; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&valorDaContagem, 1)
			v := atomic.LoadInt64(&valorDaContagem)
			fmt.Println(v)
			runtime.Gosched()
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

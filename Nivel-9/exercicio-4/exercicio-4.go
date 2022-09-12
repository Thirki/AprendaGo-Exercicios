package main

import (
	"fmt"
	"runtime"
	"sync"
)

var valorDaContagem int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	loopTimes := 1000
	wg.Add(loopTimes)
	for i := 0; i < loopTimes; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
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
// - Utilize mutex para consertar a condição de corrida do exercício anterior.

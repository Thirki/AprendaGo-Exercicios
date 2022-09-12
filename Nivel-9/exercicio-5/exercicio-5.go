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
// - Utilize atomic para consertar a condição de corrida do exercício #3.

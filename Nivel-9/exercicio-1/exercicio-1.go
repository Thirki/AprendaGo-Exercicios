package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	func() {
		defer wg.Done()
		fmt.Println("Olá")
	}()
	func() {
		defer wg.Done()
		fmt.Println("Tudo bem")
	}()
	wg.Wait()
}

// @TODO:
// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

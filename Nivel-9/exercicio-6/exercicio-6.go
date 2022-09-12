package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	arch := runtime.GOARCH
	fmt.Printf("My OS is: \t%v \n", os)
	fmt.Printf("My ARCH is: \t%v \n", arch)
}

// @TODO:
// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//     - go run
//     - go build
//     - go install

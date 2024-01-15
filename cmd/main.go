package main

import (
	"fmt"
	"github.com/baudekin/calc-ex-go/pkg"
)

func main() {
	var pos int64
	pos = 6
	fmt.Println(pkg.FibForLoop(pos))
	fmt.Println(pkg.FibRecursion(pos))
	fmt.Println(pkg.FibTailRecusion(pos))
	fmt.Println(pkg.FibChan(pos))
}

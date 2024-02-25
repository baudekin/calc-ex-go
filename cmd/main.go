package main

import (
	"fmt"
	"github.com/baudekin/calc-ex-go/pkg"
)

func main() {
	var pos uint64
	pos = 1000
	fmt.Println(pkg.FibForLoop(pos))
	fmt.Println(pkg.FibTailRecusion(pos))
	fmt.Println(pkg.FibChan(pos))
	fmt.Println(pkg.FibMath(pos))
	//fmt.Println(pkg.FibRecursion(pos))
}

package pkg

import (
	"fmt"
	"math"
	"math/big"
)

func FibForLoop(pos uint64) uint64 {
	var res, ind, first, second uint64

	for ind, first, second = 0, 0, 1; ind <= pos; ind, first, second = ind+1, first+second, first {
		if ind == pos {
			res = first
		}
	}

	return res
}

func FibRecursion(pos uint64) uint64 {
	if pos < 2 {
		return pos
	}
	return FibRecursion(pos-2) + FibRecursion(pos-1)
}

func FibTailRecusion(pos uint64) uint64 {
	return fibTail(pos, 0, 1)
}

func fibTail(pos uint64, first uint64, second uint64) uint64 {
	if pos == 0 {
		return first
	}
	return fibTail(pos-1, second, first+second)
}

func FibMath(pos uint64) uint64 {
	sqrt5Plus1 := 1 + math.Sqrt(5)
	x := math.Pow(sqrt5Plus1, float64(pos))
	z := math.Pow(2.0, float64(pos)) * math.Sqrt(5)
	fmt.Printf("x: %d, z: %d", uint64(x), uint64(z))
	res := x / z
	// Past 75th position floating point errors creep in
	return uint64(math.Round(res))
}

func FibBigMath(pos uint64) uint64 {
	// Note if you use math with uint64 this whoud fail over 75th position
	bPos := big.NewFloat(float64(pos))
	bPosF, _ := bPos.Float64()
	fmt.Printf("bPos=%f\n", bPosF)
	five := big.NewFloat(5.0)
	fiveF, _ := five.Float64()
	fmt.Printf("five=%f\n", fiveF)
	sqrtFive := five.Sqrt(five)
	sqrtFiveF, _ := sqrtFive.Float64()
	fmt.Printf("sqrtFive=%f\n", sqrtFiveF)
	sqrtFivePlusOne := sqrtFive.Add(sqrtFive, big.NewFloat(1.0))
	sqrtFivePlusOneF, _ := sqrtFivePlusOne.Float64()
	fmt.Printf("sqrtFivePlusOne=%f\n", sqrtFivePlusOneF)
	x := pow(sqrtFivePlusOne, pos)
	xF, _ := x.Float64()
	fmt.Printf("xF=%f\n", xF)
	z := pow(big.NewFloat(2.0), pos)
	z.Mul(z, new(big.Float).Sqrt(big.NewFloat(5.0)))
	zF, _ := z.Float64()
	fmt.Printf("zF=%f\n", zF)

	return 0
}

func pow(base *big.Float, exp uint64) *big.Float {
	result := new(big.Float)
	result.SetPrec(256)
	result.Set(base)
	for i := uint64(0); i < exp-1; i++ {
		result.Mul(result, base)
	}
	return result
}

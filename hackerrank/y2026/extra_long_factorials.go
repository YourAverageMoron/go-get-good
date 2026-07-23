package y2026

import (
	"fmt"
	"math/big"
)


func ExtraLongFactorials(n int32) {
    res := RecursiveFactorial(big.NewInt(1), int64(n))

    fmt.Printf("%v\n", res)
}

func RecursiveFactorial(curr *big.Int, n int64) *big.Int{
    if n == 1 {
        return curr
    }
    mul := big.NewInt(0)
    mul.Mul(curr, big.NewInt(n))
    return RecursiveFactorial(mul, n-1)
}

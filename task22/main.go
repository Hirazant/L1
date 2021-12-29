package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("456543256743256786543234567543567", 10)
	b := new(big.Int)
	b.SetString("34567865432456789765437", 10)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("res - %v", div)

	mult := new(big.Int)
	mult.Mul(a, b)
	fmt.Println("res - %v", mult)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println("res - %v", sum)

	dif := new(big.Int)
	dif.Sub(a, b)
	fmt.Println("res - %v", dif)

}

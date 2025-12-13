package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	a.SetString("3000000000", 10) //  a > 2^20

	b := new(big.Int)
	b.SetString("1500000000", 10) //  b > 2^20

	fmt.Printf("Число A: %s\n", a.String())
	fmt.Printf("Число B: %s\n\n", b.String())

	res := new(big.Int)

	// 1. Сложение (a + b)
	res.Add(a, b)
	fmt.Printf("Сложение: %s\n", res.String())

	// 2. Вычитание (a - b)
	res.Sub(a, b)
	fmt.Printf("Вычитание: %s\n", res.String())

	// 3. Умножение (a * b)
	res.Mul(a, b)
	fmt.Printf("Умножение: %s\n", res.String())

	// 4. Деление (a / b)
	res.Div(a, b)
	fmt.Printf("Деление: %s\n", res.String())
}

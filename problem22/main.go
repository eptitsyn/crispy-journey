package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
	значение которых > 2^20 = 1048576
*/
//done

func main() {
	a := big.NewInt(1048578)
	b := big.NewInt(1048577)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)
	add := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)

	fmt.Println("Mul:", mul)
	fmt.Println("Div:", div)
	fmt.Println("Add:", add)
	fmt.Println("Sub:", sub)
}

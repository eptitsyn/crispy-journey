package main

import "fmt"

/*
	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/
//done

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed := string(runes)
	return reversed
}

func main() {
	input := "главрыба"
	result := reverseString(input)
	fmt.Printf("%s - %s\n", input, result)
}

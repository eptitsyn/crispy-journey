package main

import (
	"fmt"
	"unicode"
)

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/
//done

func checkUnique(str string) bool {
	result := make(map[rune]bool)
	for _, c := range str {
		char := unicode.ToLower(c)
		if _, ok := result[char]; ok {
			return false
		}
		result[char] = true
	}
	return true
}

func main() {
	strings := []string{"abcd", "abCdefAaf", "aabcd", "aAbcd"}
	for _, str := range strings {
		fmt.Printf("%s - %t\n", str, checkUnique(str))
	}
}

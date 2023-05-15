package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

/*
Ответ: v останется в памяти и не будет собран сборщиком т.к. на него будет ссылаться justString
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
	fmt.Println(justString)
}

func createHugeString(i int) string {
	return strings.Repeat("a", i)
}

func main() {
	someFunc()
}

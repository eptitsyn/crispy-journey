package main

import "fmt"

/*
 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setOne(num int64, i int) int64 {
	return num ^ (1 << i)
}

func setZero(num int64, i int) int64 {
	return num &^ (1 << i)
}

func main() {
	var i int
	var num int64 = 65535

	fmt.Print("input i:")
	if _, err := fmt.Scanf("%d", &i); err != nil {
		panic(err)
	}
	fmt.Println(setOne(num, i))
	fmt.Println(setZero(num, i))
}

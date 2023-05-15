package main

import (
	"fmt"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
func calcSquares() {
	arr := []int{2, 4, 6, 8, 10}
	result := make(chan int)
	for _, v := range arr {
		go func(num int) {
			result <- num * num
		}(v)
	}

	//collect results through chan
	sum := 0
	for i := 0; i < 5; i++ {
		select {
		case num := <-result:
			sum += num
		}
	}
	fmt.Println(sum)
}

func main() {
	calcSquares()
}

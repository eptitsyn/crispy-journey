package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
	и выведет их квадраты в stdout.
*/
//done

func main() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}

	wg.Add(len(arr))
	for _, v := range arr {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(v)
	}
	wg.Wait()
}

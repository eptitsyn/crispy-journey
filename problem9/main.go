package main

import (
	"fmt"
	"sync"
)

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
	во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		defer wg.Done()
		for num := range chan1 {
			chan2 <- num * 2
		}
		close(chan2)
	}()
	go func() {
		defer wg.Done()
		for num := range chan2 {
			fmt.Println(num)
		}
	}()

	for i := 0; i <= 100; i++ {
		chan1 <- i
	}
	close(chan1)
	wg.Wait()
}

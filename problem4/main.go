package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
*/

func main() {
	var workerCount int
	fmt.Print("Enter worker count: ")
	if _, err := fmt.Scanf("%d", &workerCount); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	jobs := make(chan int)
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		//worker
		go func(id int, jobs <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d recieved %d\n", id, job)
			}
		}(i, jobs, &wg)
	}

	//push messages to channel
	go func() {
		for {
			jobs <- rand.Int()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	close(jobs)
}
